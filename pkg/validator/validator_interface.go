package validator

import (
	"money-tracker/pkg/database_manager"
	"money-tracker/pkg/operation"
	"net/http"
)

type ValidatorInterface interface {
	Validate(operation.Operation, *http.Request, database_manager.DatabaseManagerInterface) *ValidationResult
}
