package validator

import (
	"money-tracker/pkg/database/database_manager"
	"money-tracker/pkg/validation/validation_result"
)

type ValidatorInterface interface {
	Validate([]byte, database_manager.DatabaseManagerInterface) *validation_result.ValidationResult
}
