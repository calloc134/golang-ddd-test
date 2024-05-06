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

		ulid, err := domain.NewULID(userTable.UserULID)

		if err != nil {
			return nil, err
		}

		name, err := domain.NewName(userTable.UserDetail.Name)

		if err != nil {
			return nil, err
		}

		age, err := domain.NewAge(userTable.UserDetail.Age)

		if err != nil {
			return nil, err
		}

		users = append(users, domain.User{
			ULID:    ulid,
			Version: userTable.Version,
			UserDetail: &domain.UserDetail{
				Name: name,
				Age:  age,
			},
		})
	}

	return users, nil
}

func (ur UserRepository) FindByUlid(context context.Context, ulid domain.UlidValue) (*domain.User, error) {

	userTable := schemas.UserTable{}

	err := ur.db.NewSelect().Model(&userTable).Where("user_ulid = ?", ulid.String()).Relation("UserDetail").Scan(context)

	if err != nil {
		return nil, err
	}

	// ulid変数に再代入してるのも気になるが、同じ値になるはずだが
	ulid, err = domain.NewULID(userTable.UserULID)

	if err != nil {
		return nil, err
	}

	name, err := domain.NewName(userTable.UserDetail.Name)

	if err != nil {
		return nil, err
	}

	age, err := domain.NewAge(userTable.UserDetail.Age)

	if err != nil {
		return nil, err
	}

	return &domain.User{
		ULID: ulid,
		UserDetail: &domain.UserDetail{
			Name: name,
			Age:  age,
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
	versionCheckErr := tx.NewSelect().Model(&schemas.UserTable{}).Column("version").Where("user_ulid = ?", user.ULID).Scan(context, &version)

	// データがあってversionが一致しない場合はエラー
	if versionCheckErr != nil && version != user.Version {
		return nil, errors.New("楽観的ロックエラー")
	}

	userTable := schemas.UserTable{
		UserULID:       user.ULID.String(),
		Version:        user.Version + 1,
		UserDetailULID: user.UserDetail.ULID.String(),
	}

	userDetailTable := schemas.UserDetailTable{
		UserDetailULID: user.UserDetail.ULID.String(),
		Name:           user.UserDetail.Name.String(),
		Age:            user.UserDetail.Age.Int(),
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
		_, err = tx.NewUpdate().Model(&userTable).Set("version = ?, user_detail_ulid = ?", userTable.Version, userTable.UserDetailULID).Where("user_ulid = ?", userTable.UserULID).Exec(context)

		if err != nil {
			return nil, err
		}
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return user, nil
}
