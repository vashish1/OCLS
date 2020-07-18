package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var port = os.Getenv("PORT")
var r, router *mux.Router

func setupRoutes() {
	// r.Handle("/login",v1.login()).Methods("POST")
}

func main() {
	router = mux.NewRouter()
	r = router.PathPrefix("/api").Subrouter()
	setupRoutes()
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
