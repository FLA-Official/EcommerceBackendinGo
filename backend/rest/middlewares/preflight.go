package middlewares

import (
	"log"
	"net/http"
)

// This OPTIONS method actually meant to handle preflight request
// Preflight request is a request send by the browser to check if I gave the required permissions for the complex requests
// The purpose is for security, so if they dont find the necessary perms then the browser wont allow it to run the server and show the error 400 bad gateway.
// This globalRouter will check if there is anyb OPTION method called by browser. With this we are handling the OPTION method (preflight request) globally. The moment server launches the browser will instantly check if there is OPTION method and then proceed with other.
func Preflight(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Globally Handling preflight
		if r.Method == "OPTIONS" {
			log.Println("Handled Preflight")
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
