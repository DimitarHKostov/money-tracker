package app

import (
	"fmt"
	"net/http"

	"money-tracker/pkg/api"
	"money-tracker/pkg/database/database_config"
	"money-tracker/pkg/database/database_credentials_formatter_factory"
	"money-tracker/pkg/database/database_manager_factory"

	"money-tracker/pkg/operation"
	"money-tracker/pkg/validation/validator_factory"

	"money-tracker/pkg/jwt"

	"github.com/gorilla/mux"
)

type App struct {
	AppRouter      *mux.Router
	JWTManager     jwt.JWTManagerInterface
	DatabaseConfig database_config.DatabaseConfig
	AppVersion     string
	AppPort        string
}

func (a *App) Run() error {
	databaseCredentialsFormatterFactory := &database_credentials_formatter_factory.DatabaseCredentialsFormatterFactory{}
	dbConnectionString := databaseCredentialsFormatterFactory.ProduceFormattedCredentials(a.DatabaseConfig)

	databaseManagerFactory := &database_manager_factory.DatabaseManagerFactory{}
	databaseManager, err := databaseManagerFactory.ProduceDatabaseManager(a.DatabaseConfig.DatabaseType, dbConnectionString)
	if err != nil {
		return err
	}

	validatorFactory := &validator_factory.ValidatorFactory{}
	registerValidator := validatorFactory.ProduceValidator(operation.Register)
	loginValidator := validatorFactory.ProduceValidator(operation.Login)

	path := fmt.Sprintf("/api/%s", a.AppVersion)

	a.AppRouter.PathPrefix(path+"/register").HandlerFunc(api.Register(registerValidator, databaseManager)).Methods(http.MethodPost, http.MethodOptions)
	a.AppRouter.PathPrefix(path+"/login").HandlerFunc(api.Login(loginValidator, databaseManager)).Methods(http.MethodPost, http.MethodOptions)

	return http.ListenAndServe(a.AppPort, a.AppRouter)
}
