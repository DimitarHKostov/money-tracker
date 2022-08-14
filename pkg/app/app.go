package app

import (
	"net/http"

	"money-tracker/pkg/api"
	"money-tracker/pkg/database_manager"
	"money-tracker/pkg/jwt"
	"money-tracker/pkg/validator"

	"github.com/gorilla/mux"
)

const (
	apiVersion = "/api/v1"
	appPort    = ":80"
)

type App struct {
	Router          *mux.Router
	Validator       validator.ValidatorInterface
	JWTManager      jwt.JWTManagerInterface
	DatabaseManager database_manager.DatabaseManagerInterface
}

func (a *App) Run() error {
	a.Router.PathPrefix(apiVersion+"/register").HandlerFunc(api.Register(a.Validator, a.DatabaseManager)).Methods(http.MethodPost, http.MethodOptions)
	//a.Router.PathPrefix(apiVersion+"/login").HandlerFunc(api.Login(a.Validator, a.DbConnection, a.Query, a.JWTManager)).Methods(http.MethodPost, http.MethodOptions)
	//a.Router.PathPrefix(apiVersion+"/logout").HandlerFunc(api.Logout(a.Validator, a.DbConnection, a.Query)).Methods(http.MethodPost, http.MethodOptions)
	//a.Router.PathPrefix(apiVersion+"/refresh").HandlerFunc(api.Refresh(a.Validator, a.DbConnection, a.Query)).Methods(http.MethodPost, http.MethodOptions)
	//a.Router.PathPrefix(apiVersion+"/calculate").HandlerFunc(api.Calculate(a.Validator, a.DbConnection, a.Query)).Methods(http.MethodPost, http.MethodOptions)

	return http.ListenAndServe(appPort, a.Router)
}
