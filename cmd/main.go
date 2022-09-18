package main

import (
	"log"
	"money-tracker/pkg/app"
	"money-tracker/pkg/cors"
	"money-tracker/pkg/database/database_config"
	"money-tracker/pkg/generator"
	"money-tracker/pkg/jwt"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

const (
	apiVersion = "v1"
	appPort    = ":80"
	secretKey  = "asdasasdasasdasasdasasdasaa"
	host       = "localhost"
	port       = "5432"
	user       = "postgres"
	password   = "test"
)

func main() {
	app := app.App{
		AppRouter: mux.NewRouter(),
		JWTManager: &jwt.JWTManager{
			PayloadGenerator: &generator.PayloadGenerator{},
			SecretKey:        secretKey,
		},
		DatabaseConfig: database_config.DatabaseConfig{
			Host:     host,
			Port:     port,
			Username: user,
			Password: password,
		},
		AppVersion: apiVersion,
		AppPort:    appPort,
	}

	app.AppRouter.Use(cors.Middleware)

	err := app.Run()
	if err != nil {
		log.Panic(err)
	}
}
