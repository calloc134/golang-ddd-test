package repositories

import (
	"context"

	"github.com/calloc134/golang-ddd-test/application"
	"github.com/calloc134/golang-ddd-test/cmd/migrate/schemas"
	"github.com/calloc134/golang-ddd-test/domain"
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
		})
	}

	return users, nil
}

func (ur UserRepository) FindByID(context context.Context, uuid string) (domain.UserAggregate, error) {
	
	userTable := schemas.UserTable{}

	err := ur.db.NewSelect().Model(&userTable).Where("id = ?", uuid).Scan(context)

	if err != nil {
		return domain.UserAggregate{}, err
	}

	return domain.UserAggregate{
		ULID: userTable.ULID,
		Name: userTable.Name,
		Age: userTable.Age,
	}, nil
	
	
}

func (ur UserRepository) Save(context context.Context, user domain.UserAggregate) error {
	
	userTable := schemas.UserTable{
		ULID: user.ULID,
		Name: user.Name,
		Age: user.Age,
	}

	_, err := ur.db.NewInsert().Model(&userTable).Exec(context)

	if err != nil {
		return err
	}

	return nil
}


