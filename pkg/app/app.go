package app

import (
	"database/sql"
	"net/http"

	"money-tracker/pkg/api"
	_ "money-tracker/pkg/api"

	"github.com/gorilla/mux"
)

const (
	basePath = "/api"
	appPort  = ":80"
)

type App struct {
	Router       *mux.Router
	DbConnection *sql.DB
}

func (a *App) Run() error {
	return http.ListenAndServe(appPort, a.Router)
}

func (a *App) Configure() error {
	a.configurePaths(a.DbConnection)

	return nil
}

func (a *App) configurePaths(dbConnection *sql.DB) {
	a.Router.PathPrefix(basePath+"/register").HandlerFunc(a.configureOperation(api.Register)).Methods(http.MethodPost, http.MethodOptions)
	a.Router.PathPrefix(basePath+"/login").HandlerFunc(a.configureOperation(api.Login)).Methods(http.MethodPost, http.MethodOptions)
}

func (a *App) configureOperation(operation func(*sql.DB) func(w http.ResponseWriter, r *http.Request)) func(http.ResponseWriter, *http.Request) {
	return operation(a.DbConnection)
}
