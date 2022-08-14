package app

import (
	"net/http"

	"money-tracker/pkg/api"
	"money-tracker/pkg/database_manager"
	"money-tracker/pkg/jwt"
	"money-tracker/pkg/operation"
	"money-tracker/pkg/validation/validator_factory"

	"github.com/gorilla/mux"
)

const (
	apiVersion = "/api/v1"
	appPort    = ":80"
)

type App struct {
	Router           *mux.Router
	ValidatorFactory validator_factory.ValidatorFactory
	JWTManager       jwt.JWTManagerInterface
	DatabaseManager  database_manager.DatabaseManagerInterface
}

func (a *App) Run() error {
	a.Router.PathPrefix(apiVersion+"/register").HandlerFunc(api.Register(a.ValidatorFactory.ProduceValidator(operation.Register), a.DatabaseManager)).Methods(http.MethodPost, http.MethodOptions)

	return http.ListenAndServe(appPort, a.Router)
}
