package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/calloc134/golang-ddd-test/cmd/api/graph"
	"github.com/calloc134/golang-ddd-test/src/mutation/application"
	"github.com/calloc134/golang-ddd-test/src/mutation/repository"
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
		bundebug.WithEnabled(true),
		bundebug.FromEnv(""),
	))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(
		application.NewUserApplication(repository.NewUserRepository(db)),
		application.NewPostApplication(repository.NewPostRepository(db)),
	)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
