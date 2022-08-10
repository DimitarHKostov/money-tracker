package api

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"money-tracker/types"
	"net/http"
)

const (
	registerQuery = `
	INSERT INTO Account(Email, Username, Password)
	VALUES ($1, $2, $3)`
)

func Register(dbConnection *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Fatalln(err)
		}

		var account types.Account
		json.Unmarshal(body, &account)

		_, err = dbConnection.Exec(registerQuery, account.Email, account.Username, account.Password)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
