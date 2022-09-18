package api

import (
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	a, err := r.Cookie("asd")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if a != nil {
		log.Println((*a).String())
	}

	w.WriteHeader(http.StatusOK)
}
