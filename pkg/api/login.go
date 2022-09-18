package api

import (
	"io"
	"log"
	"money-tracker/pkg/database/database_manager"
	"money-tracker/pkg/validation/validation_result_status"
	"money-tracker/pkg/validation/validator"
	"net/http"
)

func Login(Validator validator.ValidatorInterface, databaseManager database_manager.DatabaseManagerInterface) func(w http.ResponseWriter, r *http.Request) {
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

		validationResult := Validator.Validate(body, databaseManager)
		if validationResult.ValidationResultStatus == validation_result_status.Failure {
			log.Println(validationResult.ValidationResultMessage)
			w.WriteHeader(http.StatusForbidden)
			return
		}

		// _, err = databaseManager.RegisterAccount(account)
		// if err != nil {
		// 	log.Println(err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		w.WriteHeader(http.StatusOK)
	}
}
