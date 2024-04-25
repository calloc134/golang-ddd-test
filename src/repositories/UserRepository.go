package repositories

import (
	"context"
	"errors"

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

func (ur UserRepository) FindAll(context context.Context) ([]domain.UserAggregate, error) {

	userTables := []schemas.UserTable{}


	err := ur.db.NewSelect().Model((&userTables)).Scan(context)

	if err != nil {
		return nil, err
	}

	var users []domain.UserAggregate
	
	for _, userTable := range userTables {
		users = append(users, domain.UserAggregate{
			ULID: userTable.ULID,
			Name: userTable.Name,
			Age: userTable.Age,
			Version: userTable.Version,
		})
	}

	return users, nil
}

func (ur UserRepository) FindByUlid(context context.Context, uuid string) (domain.UserAggregate, error) {
	
	userTable := schemas.UserTable{}

	err := ur.db.NewSelect().Model(&userTable).Where("id = ?", uuid).Scan(context)

	if err != nil {
		return domain.UserAggregate{}, err
	}

	return domain.UserAggregate{
		ULID: userTable.ULID,
		Name: userTable.Name,
		Age: userTable.Age,
		Version: userTable.Version,
	}, nil
	
	
}

func (ur UserRepository) Save(context context.Context, user domain.UserAggregate) error {

	tx, err := ur.db.BeginTx(context, nil)

	if err != nil {
		return err
	}

	defer tx.Rollback()

	// 楽観的ロックの判定
	exists, err := tx.NewSelect().Model(&schemas.UserTable{}).Where("ulid = ? and version = ?", user.ULID, user.Version).Exists(context)

	if err != nil {
		return err
	}

	if !exists {
		return errors.New("楽観的ロックエラー")
	}

	userTable := schemas.UserTable{
		ULID: user.ULID,
		Name: user.Name,
		Age: user.Age,
		Version: user.Version + 1,
	}

	_, err = tx.NewInsert().Model(&userTable).Exec(context);

	if err != nil {
		return err
	}

	err = tx.Commit()

	if err != nil {
		return err
	}


	return nil
}


