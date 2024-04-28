package repositories

import (
	"context"
	"errors"
	"fmt"

	"github.com/calloc134/golang-ddd-test/cmd/migrate/schemas"
	"github.com/calloc134/golang-ddd-test/src/application"
	"github.com/calloc134/golang-ddd-test/src/domain"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func NewUserRepository(db *bun.DB) application.IUserRepository {
	return UserRepository{db: db}
}

func (ur UserRepository) FindAll(context context.Context) ([]domain.User, error) {

	userTables := []schemas.UserTable{}

	err := ur.db.NewSelect().Model((&userTables)).Relation("UserDetail").Scan(context)

	if err != nil {
		return nil, err
	}

	var users []domain.User

	for _, userTable := range userTables {
		users = append(users, domain.User{
			ULID:    userTable.ULID,
			Version: userTable.Version,
			UserDetail: &domain.UserDetail{
				Name: userTable.UserDetail.Name,
				Age:  userTable.UserDetail.Age,
			},
		})
	}

	return users, nil
}

func (ur UserRepository) FindByUlid(context context.Context, uuid string) (*domain.User, error) {

	userTable := schemas.UserTable{}

	err := ur.db.NewSelect().Model(&userTable).Where("ulid = ?", uuid).Relation("UserDetail").Scan(context)

	if err != nil {
		return nil, err
	}

	return &domain.User{
		ULID: userTable.ULID,
		UserDetail: &domain.UserDetail{
			Name: userTable.UserDetail.Name,
			Age:  userTable.UserDetail.Age,
		},
		Version: userTable.Version,
	}, nil

}

func (ur UserRepository) Save(context context.Context, user *domain.User) (*domain.User, error) {

	tx, err := ur.db.BeginTx(context, nil)

	if err != nil {
		return nil, err
	}

	defer tx.Rollback()

	// 楽観的ロックの判定
	// ulidが一致したらversionを取得 versionが一致したらパス
	var version int
	versionCheckErr := tx.NewSelect().Model(&schemas.UserTable{}).Column("version").Where("ulid = ?", user.ULID).Scan(context, &version)

	// データがあってversionが一致しない場合はエラー
	if versionCheckErr != nil && version != user.Version {
		return nil, errors.New("楽観的ロックエラー")
	}

	userTable := schemas.UserTable{
		ULID:    user.ULID,
		Version: user.Version + 1,
	}

	userDetailTable := schemas.UserDetailTable{
		Name:   user.UserDetail.Name,
		Age:    user.UserDetail.Age,
		UserID: user.ULID,
	}

	_, err = tx.NewInsert().Model(&userDetailTable).Exec(context)

	if err != nil {
		return nil, err
	}

	if versionCheckErr != nil && versionCheckErr.Error() == "sql: no rows in result set" {
		fmt.Println("insert")
		_, err = tx.NewInsert().Model(&userTable).Exec(context)

		if err != nil {
			return nil, err
		}
	} else {
		fmt.Println("update")
		_, err = tx.NewUpdate().Model(&userTable).Set("version = ?", userTable.Version).Where("ulid = ?", user.ULID).Exec(context)

		if err != nil {
			return nil, err
		}
	}

	err = tx.Commit()

	if err != nil {
		return nil, err
	}

	return user, nil
}
