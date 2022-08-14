package main

import (
	"database/sql"
	"fmt"
	"log"
	"money-tracker/pkg/app"
	"money-tracker/pkg/cors"
	"money-tracker/pkg/database_manager"
	"money-tracker/pkg/generator"
	"money-tracker/pkg/jwt"
	"money-tracker/pkg/query"
	"money-tracker/pkg/validator"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "test"
)

var (
	databaseConnectionString = fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
)

func main() {
	dbConnection, err := sql.Open("postgres", databaseConnectionString)
	if err != nil {
		log.Panic(err)
	}
	defer dbConnection.Close()

	app := app.App{
		Router: mux.NewRouter(),
		DatabaseManager: &database_manager.DatabaseManager{
			DbConnection: dbConnection,
			Query:        &query.Query{},
		},
		Validator: &validator.Validator{},
		JWTManager: &jwt.JWTManager{
			PayloadGenerator: &generator.PayloadGenerator{},
			SecretKey:        "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		}}
	app.Router.Use(cors.Middleware)

	err = app.Run()
	if err != nil {
		log.Panic(err)
	}
}
