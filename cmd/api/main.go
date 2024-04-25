package main

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {

	sqldb, err := sql.Open(sqliteshim.DriverName(), "file:test.s3db?cache=shared")

	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithEnabled(false),
		bundebug.FromEnv(""),
	))

	// ctx := context.Background()

	// application := application.NewUserApplication(repositories.NewUserRepository(db))

	// ec := echo.New()

}