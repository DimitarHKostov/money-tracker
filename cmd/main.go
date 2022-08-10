package main

import (
	"database/sql"
	"fmt"
	"log"
	"money-tracker/pkg/app"
	"money-tracker/pkg/cors"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s sslmode=disable",
		host, port, user, password)

	dbConnection, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer dbConnection.Close()

	app := app.App{Router: mux.NewRouter(), DbConnection: dbConnection}
	app.Router.Use(cors.Middleware)

	app.Configure()
	err = app.Run()
	if err != nil {
		log.Panic(err)
	}
}
