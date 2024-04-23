package main

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/calloc134/golang-ddd-test/application"
	"github.com/calloc134/golang-ddd-test/repositories"
	"github.com/labstack/echo/v4"
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

	ctx := context.Background()

	application := application.NewUserApplication(repositories.NewUserRepository(db))

	ec := echo.New()



	ec.GET("/GetAllUser", func(c echo.Context) error {
		users, err := application.FindAll(ctx)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, users)
	})


	ec.GET("/GetUserByID", func(c echo.Context) error {
		users, err := application.FindByID(ctx, "test")
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		return c.JSON(http.StatusOK, users)
	})

	// コントローラとしてのフレームワークは後で追加

	// http.HandleFunc("/NewUser", func(w http.ResponseWriter, r *http.Request) {
	// 	err := application.Save(ctx, "test")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	fmt.Fprintf(w, "User created")http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// })


	ec.Start(":1323")
}