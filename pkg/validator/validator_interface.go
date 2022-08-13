package validator

import (
	"database/sql"
	"money-tracker/pkg/operation"
	"money-tracker/pkg/query"
	"net/http"
)

type ValidatorInterface interface {
	Validate(operation.Operation, *http.Request, *sql.DB, query.QueryInterface) *ValidationResult
}
