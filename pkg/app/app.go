package app

import (
	"net/http"

	"money-tracker/pkg/api"
	"money-tracker/pkg/database/database_manager_factory"
	"money-tracker/pkg/database/database_type"

	"money-tracker/pkg/operation"
	"money-tracker/pkg/validation/validator_factory"

	"money-tracker/pkg/jwt"

	"github.com/gorilla/mux"
)

const (
	apiVersion = "/api/v1"
	appPort    = ":80"
)

type App struct {
	Router       *mux.Router
	JWTManager   jwt.JWTManagerInterface
	DatabaseType database_type.DatabaseType
}

func (a *App) Run() error {
	databaseManagerFactory := &database_manager_factory.DatabaseManagerFactory{}
	databaseManager, err := databaseManagerFactory.ProduceDatabaseManager(a.DatabaseType)
	if err != nil {
		return err
	}

	validatorFactory := &validator_factory.ValidatorFactory{}
	registerValidator := validatorFactory.ProduceValidator(operation.Register)

	a.Router.PathPrefix(apiVersion+"/register").HandlerFunc(api.Register(registerValidator, databaseManager)).Methods(http.MethodPost, http.MethodOptions)

	return http.ListenAndServe(appPort, a.Router)
}
