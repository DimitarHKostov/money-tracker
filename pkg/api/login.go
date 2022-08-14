package api

import (
	"database/sql"
	"money-tracker/pkg/jwt"
	"money-tracker/pkg/query"
	"money-tracker/pkg/validator"
	"net/http"
)

func Login(validatorInstance validator.ValidatorInterface, dbConnection *sql.DB, queryInstance query.QueryInterface, JWTManagerInstance jwt.JWTManagerInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}
