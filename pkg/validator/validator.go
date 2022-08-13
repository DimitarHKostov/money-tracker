package validator

import (
	"database/sql"
	"money-tracker/pkg/operation"
	"money-tracker/pkg/query"
	"net/http"
)

type Validator struct{}

func (v *Validator) Validate(op operation.Operation, request *http.Request, dbConnection *sql.DB, queryInstance query.QueryInterface) *ValidationResult {
	switch op {
	case operation.Register:
		return v.validateRegister(request, dbConnection, queryInstance)
	case operation.Login:
		return v.validateLogin(request, dbConnection, queryInstance)
	case operation.Logout:
		return v.validateLogout(request, dbConnection, queryInstance)
	case operation.Refresh:
		return v.validateRefresh(request, dbConnection, queryInstance)
	case operation.Calculate:
		return v.validateCalculate(request, dbConnection, queryInstance)
	default:
		return nil
	}
}

func (v *Validator) validateRegister(request *http.Request, dbConnection *sql.DB, queryInstance query.QueryInterface) *ValidationResult {
	return &ValidationResult{
		ValidationResultStatus: Success,
		ValidationResultMessage: &ValidationResultMessage{
			Message: Success.String(),
		},
	}
}

func (v *Validator) validateLogin(request *http.Request, dbConnection *sql.DB, queryInstance query.QueryInterface) *ValidationResult {
	return &ValidationResult{
		ValidationResultStatus: Success,
		ValidationResultMessage: &ValidationResultMessage{
			Message: Success.String(),
		},
	}
}

func (v *Validator) validateLogout(request *http.Request, dbConnection *sql.DB, queryInstance query.QueryInterface) *ValidationResult {
	return &ValidationResult{
		ValidationResultStatus: Success,
		ValidationResultMessage: &ValidationResultMessage{
			Message: Success.String(),
		},
	}
}

func (v *Validator) validateRefresh(request *http.Request, dbConnection *sql.DB, queryInstance query.QueryInterface) *ValidationResult {
	return &ValidationResult{
		ValidationResultStatus: Success,
		ValidationResultMessage: &ValidationResultMessage{
			Message: Success.String(),
		},
	}
}

func (v *Validator) validateCalculate(request *http.Request, dbConnection *sql.DB, queryInstance query.QueryInterface) *ValidationResult {
	return &ValidationResult{
		ValidationResultStatus: Success,
		ValidationResultMessage: &ValidationResultMessage{
			Message: Success.String(),
		},
	}
}
