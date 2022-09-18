package api

import (
	"encoding/json"
	"io"
	"log"
	"money-tracker/pkg/database/database_manager"
	"money-tracker/pkg/jwt"
	"money-tracker/pkg/validation/validation_result_status"
	"money-tracker/pkg/validation/validator"
	"money-tracker/types"
	"net/http"
	"time"
)

const (
	sessionCookie = "sessionCookie"
)

func Login(validator *validator.ValidatorInterface, databaseManager *database_manager.DatabaseManagerInterface, JWTManager jwt.JWTManagerInterface) func(w http.ResponseWriter, r *http.Request) {
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

		token, err := JWTManager.GenerateToken(&account, 5*time.Minute)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		c := http.Cookie{Name: "asd", Value: token, Expires: time.Now().Add(5 * time.Minute), HttpOnly: true}

		http.SetCookie(w, &c)

		w.WriteHeader(http.StatusOK)
	}
}
