package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"money-tracker/pkg/operation"
	"money-tracker/pkg/query"
	"money-tracker/pkg/validator"
	"money-tracker/types"
)

func Register(validatorInstance validator.ValidatorInterface, dbConnection *sql.DB, queryInstance query.QueryInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			return
		}

		validationResult := validatorInstance.Validate(operation.Register, r, dbConnection, queryInstance)
		if validationResult.ValidationResultStatus == validator.Failure {
			log.Println("validation failed", validationResult.ValidationResultMessage)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var account types.Account
		err = json.Unmarshal(body, &account)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = dbConnection.Exec(queryInstance.RegisterAccount(), account.Email, account.Username, account.Password)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
