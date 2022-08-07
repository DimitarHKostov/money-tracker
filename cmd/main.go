package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("success"))
	w.WriteHeader(http.StatusOK)
}

func handleHttp() {
	router := mux.NewRouter()
	router.HandleFunc("/get", handleGet).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8083", router))
}

func main() {
	handleHttp()
}
