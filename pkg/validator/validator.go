package validator

import (
	"money-tracker/pkg/database_manager"
	"money-tracker/pkg/operation"
	"net/http"
)

type Validator struct{}

func (v *Validator) Validate(op operation.Operation, request *http.Request, databaseManager database_manager.DatabaseManagerInterface) *ValidationResult {
	switch op {
	case operation.Register:
		return v.validateRegister(request, databaseManager)
	case operation.Login:
		return v.validateLogin(request, databaseManager)
	case operation.Logout:
		return v.validateLogout(request, databaseManager)
	case operation.Refresh:
		return v.validateRefresh(request, databaseManager)
	case operation.Calculate:
		return v.validateCalculate(request, databaseManager)
	default:
		return nil
	}
}

func (v *Validator) validateRegister(request *http.Request, databaseManager database_manager.DatabaseManagerInterface) *ValidationResult {
	return &ValidationResult{
		ValidationResultStatus: Success,
		ValidationResultMessage: &ValidationResultMessage{
			Message: Success.String(),
		},
	}
}

func (v *Validator) validateLogin(request *http.Request, databaseManager database_manager.DatabaseManagerInterface) *ValidationResult {
	return &ValidationResult{
		ValidationResultStatus: Success,
		ValidationResultMessage: &ValidationResultMessage{
			Message: Success.String(),
		},
	}
}

func (v *Validator) validateLogout(request *http.Request, databaseManager database_manager.DatabaseManagerInterface) *ValidationResult {
	return &ValidationResult{
		ValidationResultStatus: Success,
		ValidationResultMessage: &ValidationResultMessage{
			Message: Success.String(),
		},
	}
}

func (v *Validator) validateRefresh(request *http.Request, databaseManager database_manager.DatabaseManagerInterface) *ValidationResult {
	return &ValidationResult{
		ValidationResultStatus: Success,
		ValidationResultMessage: &ValidationResultMessage{
			Message: Success.String(),
		},
	}
}

func (v *Validator) validateCalculate(request *http.Request, databaseManager database_manager.DatabaseManagerInterface) *ValidationResult {
	return &ValidationResult{
		ValidationResultStatus: Success,
		ValidationResultMessage: &ValidationResultMessage{
			Message: Success.String(),
		},
	}
}
