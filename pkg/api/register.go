package api

import (
	"encoding/json"
	"io"
	"log"
	"money-tracker/pkg/database/database_manager"
	"money-tracker/pkg/validation/validation_result_status"
	"money-tracker/pkg/validation/validator"
	"money-tracker/types"
	"net/http"
)

func Register(validator *validator.ValidatorInterface, databaseManager *database_manager.DatabaseManagerInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		validationResult := (*validator).Validate(body, *databaseManager)
		if validationResult.ValidationResultStatus == validation_result_status.Failure {
			log.Println(validationResult.ValidationResultMessage)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		var account types.Account
		err = json.Unmarshal(body, &account)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = (*databaseManager).RegisterAccount(account)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
