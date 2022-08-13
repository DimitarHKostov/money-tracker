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
