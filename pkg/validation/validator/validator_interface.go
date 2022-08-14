package validator

import (
	"money-tracker/pkg/database_manager"
	"money-tracker/pkg/validation/validation_result"
	"net/http"
)

type ValidatorInterface interface {
	Validate(*http.Request, database_manager.DatabaseManagerInterface) *validation_result.ValidationResult
}
