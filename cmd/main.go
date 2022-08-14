package main

import (
	"log"
	"money-tracker/pkg/app"
	"money-tracker/pkg/cors"
	"money-tracker/pkg/database/database_type"
	"money-tracker/pkg/generator"
	"money-tracker/pkg/jwt"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	app := app.App{
		Router: mux.NewRouter(),
		JWTManager: &jwt.JWTManager{
			PayloadGenerator: &generator.PayloadGenerator{},
			SecretKey:        "asdasasdasasdasasdasasdasaa",
		},
		DatabaseType: database_type.Postgres}
	app.Router.Use(cors.Middleware)

	err := app.Run()
	if err != nil {
		log.Panic(err)
	}
}
