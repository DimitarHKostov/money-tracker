package handler_func_factory

import (
	"money-tracker/pkg/api"
	"money-tracker/pkg/database/database_manager"
	"money-tracker/pkg/operation"
	"money-tracker/pkg/validation/validator_factory"
	"net/http"
)

type HandlerFuncFactory struct {
	ValidatorFactory validator_factory.ValidatorFactory
}

func (hff *HandlerFuncFactory) Produce(op operation.Operation, databaseManager *database_manager.DatabaseManagerInterface) func(http.ResponseWriter, *http.Request) {
	switch op {
	case operation.Register:
		return hff.ProduceRegisterHandlerFunc(databaseManager)
	case operation.Login:
		return hff.ProduceLoginHandlerFunc(databaseManager)
	default:
		return nil
	}
}

func (hff *HandlerFuncFactory) ProduceRegisterHandlerFunc(databaseManager *database_manager.DatabaseManagerInterface) func(http.ResponseWriter, *http.Request) {
	validator := hff.ValidatorFactory.ProduceValidator(operation.Register)

	return http.HandlerFunc(api.Register(validator, *databaseManager))
}

func (hff *HandlerFuncFactory) ProduceLoginHandlerFunc(databaseManager *database_manager.DatabaseManagerInterface) func(http.ResponseWriter, *http.Request) {
	validator := hff.ValidatorFactory.ProduceValidator(operation.Login)

	return http.HandlerFunc(api.Login(validator, *databaseManager))
}
