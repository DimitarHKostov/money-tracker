package cors

import (
	"net/http"
)

func Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Methods", "POST")
		w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
		handler.ServeHTTP(w, r)
	})
}
