package app

import (
	"database/sql"
	"net/http"

	"money-tracker/pkg/api"
	"money-tracker/pkg/operation"
	"money-tracker/pkg/query"
	"money-tracker/pkg/validator"

	"github.com/gorilla/mux"
)

const (
	apiVersion = "/api/v1"
	appPort    = ":80"
)

type App struct {
	Router       *mux.Router
	DbConnection *sql.DB
	Validator    validator.ValidatorInterface
	Query        query.QueryInterface
}

func (a *App) Run() error {
	return http.ListenAndServe(appPort, a.Router)
}

func (a *App) Configure() {
	a.Router.PathPrefix(apiVersion+"/register").HandlerFunc(a.configureOperation(operation.Register)).Methods(http.MethodPost, http.MethodOptions)
}

func (a *App) configureOperation(op operation.Operation) func(http.ResponseWriter, *http.Request) {
	switch op {
	case operation.Register:
		return api.Register(a.Validator, a.DbConnection, a.Query)
	}

	return nil
}
