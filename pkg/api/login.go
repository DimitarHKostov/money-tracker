package api

import (
	"database/sql"
	"money-tracker/pkg/query"
	"money-tracker/pkg/validator"
	"net/http"
)

func Login(dbConnection *sql.DB, validator validator.ValidatorInterface, query query.QueryInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// _, err := dbConnection.Exec(registerQuery, "asd", "asd", "asd")
		// if err != nil {
		// 	log.Println(err)
		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		w.WriteHeader(http.StatusCreated)
	}
}
