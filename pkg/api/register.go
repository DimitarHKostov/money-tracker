package api

import (
	"database/sql"
	"log"
	"net/http"
)

const (
	registerQuery = `
	INSERT INTO Account(Email, Username, Password)
	VALUES ($1, $2, $3)`
)

func Register(dbConnection *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := dbConnection.Exec(registerQuery, "asd", "asd", "asd")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}
