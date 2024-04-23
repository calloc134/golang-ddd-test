package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/calloc134/golang-ddd-test/application"
	"github.com/calloc134/golang-ddd-test/repositories"
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// とりあえずテスト実装
	// コントローラとしてのフレームワークは後で追加

	// http.HandleFunc("/NewUser", func(w http.ResponseWriter, r *http.Request) {
	// 	err := application.Save(ctx, "test")
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	fmt.Fprintf(w, "User created")
	// })

	http.HandleFunc("/GetAllUser", func(w http.ResponseWriter, r *http.Request) {
		users, err := application.FindAll(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "Users: %v", users)
	})

	http.HandleFunc("/GetUserByID", func(w http.ResponseWriter, r *http.Request) {
		users, err := application.FindByID(ctx, "test")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "User: %v", users)
	})


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}