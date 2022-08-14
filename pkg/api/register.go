package api

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"money-tracker/pkg/database_manager"
	"money-tracker/pkg/operation"
	"money-tracker/pkg/validator"
	"money-tracker/types"
)

func Register(validatorInstance validator.ValidatorInterface, databaseManager database_manager.DatabaseManagerInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			return
		}

		validationResult := validatorInstance.Validate(operation.Register, r, databaseManager)
		if validationResult.ValidationResultStatus == validator.Failure {
			log.Println(validationResult.ValidationResultMessage)
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

		databaseManager.RegisterAccount(account)

		w.WriteHeader(http.StatusCreated)
	}
}
