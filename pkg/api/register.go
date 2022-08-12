package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
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
			log.Println(err)
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

		hasCollision, err := checkAccountCollisionByUsername(account.Username, dbConnection, query)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if hasCollision {
			log.Println(fmt.Sprintf("account with username %s already exists", account.Username))
			w.WriteHeader(http.StatusConflict)
			return
		}

		_, err = dbConnection.Exec(query.GetRegisterAccountQuery(), account.Email, account.Username, account.Password)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func checkAccountCollisionByUsername(username string, dbConnection *sql.DB, query query.QueryInterface) (bool, error) {
	var accountInDatabase types.Account
	err := dbConnection.QueryRow(query.GetSelectAccountWithUsernameQuery(), username).Scan(&accountInDatabase)
	if err != nil {
		log.Println(err)
		return false, err
	}

	if accountInDatabase.Username != "" {
		log.Println(fmt.Sprintf("user with username {%s} already exists in the database", accountInDatabase.Username))
		return true, nil
	}

	return false, nil
}
