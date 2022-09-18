package app

import (
	"fmt"
	"net/http"

	"money-tracker/pkg/api"
	"money-tracker/pkg/app/handler_func_factory"
	"money-tracker/pkg/database/database_config"
	"money-tracker/pkg/database/database_manager_factory"
	"money-tracker/pkg/operation"

	"money-tracker/pkg/validation/validator_factory"

	"github.com/gorilla/mux"
)

type App struct {
	AppRouter      *mux.Router
	DatabaseConfig database_config.DatabaseConfig
	AppVersion     string
	AppPort        string
}

func (a *App) Run() error {
	databaseManagerFactory := &database_manager_factory.DatabaseManagerFactory{}
	databaseManager, err := databaseManagerFactory.ProduceDatabaseManager(&a.DatabaseConfig)
	if err != nil {
		return err
	}

	handlerFuncFactory := &handler_func_factory.HandlerFuncFactory{ValidatorFactory: validator_factory.ValidatorFactory{}}
	registerHandlerFunc := handlerFuncFactory.Produce(operation.Register, &databaseManager)
	loginHandlerFunc := handlerFuncFactory.Produce(operation.Login, &databaseManager)

	path := fmt.Sprintf("/api/%s", a.AppVersion)

	a.AppRouter.PathPrefix(path+"/register").HandlerFunc(registerHandlerFunc).Methods(http.MethodPost, http.MethodOptions)
	a.AppRouter.PathPrefix(path+"/login").HandlerFunc(loginHandlerFunc).Methods(http.MethodPost, http.MethodOptions)
	a.AppRouter.PathPrefix(path+"/test").HandlerFunc(api.Test).Methods(http.MethodOptions, http.MethodGet)

	return http.ListenAndServe(a.AppPort, a.AppRouter)
}
