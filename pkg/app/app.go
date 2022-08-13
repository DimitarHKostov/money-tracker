package app

import (
	"database/sql"
	"net/http"

	"money-tracker/pkg/api"
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
	a.Router.PathPrefix(apiVersion+"/register").HandlerFunc(api.Register(a.Validator, a.DbConnection, a.Query)).Methods(http.MethodPost, http.MethodOptions)
	a.Router.PathPrefix(apiVersion+"/login").HandlerFunc(api.Login(a.Validator, a.DbConnection, a.Query)).Methods(http.MethodPost, http.MethodOptions)
	a.Router.PathPrefix(apiVersion+"/logout").HandlerFunc(api.Logout(a.Validator, a.DbConnection, a.Query)).Methods(http.MethodPost, http.MethodOptions)
	a.Router.PathPrefix(apiVersion+"/refresh").HandlerFunc(api.Refresh(a.Validator, a.DbConnection, a.Query)).Methods(http.MethodPost, http.MethodOptions)
	a.Router.PathPrefix(apiVersion+"/calculate").HandlerFunc(api.Calculate(a.Validator, a.DbConnection, a.Query)).Methods(http.MethodPost, http.MethodOptions)
}
