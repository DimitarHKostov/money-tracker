package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"money-tracker/pkg/query"
	"money-tracker/pkg/validator"
	"money-tracker/types"
	"net/http"
)

func Register(dbConnection *sql.DB, validator validator.ValidatorInterface, query query.QueryInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodOptions {
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var account types.Account
		json.Unmarshal(body, &account)

		if err := validator.ValidateUsername(account.Username); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusNotAcceptable)
			return
		}

		_, err = dbConnection.Exec(query.GetRegisterQuery(), account.Email, account.Username, account.Password)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
