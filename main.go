package main

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/vashish1/OnlineClassPortal/api/v1"
)

var port = os.Getenv("PORT")
var r, router *mux.Router

func setupRoutes() {
	r.Handle("/login", v1.StudentLogin()).Methods("POST")
	r.Handle("/check", v1.CkeckLogin()).Methods("GET")
}

func main() {
	router = mux.NewRouter()
	r = router.PathPrefix("/api").Subrouter()
	setupRoutes()
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
