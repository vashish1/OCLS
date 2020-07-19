package main

import (
	"net/http"
	"os"
    
	"github.com/gorilla/mux"
	"github.com/vashish1/OnlineClassPortal/api/v1"
	mw "github.com/vashish1/OnlineClassPortal/api/middleware"
)

var port = os.Getenv("PORT")
var r, router *mux.Router
var mwFunc mux.MiddlewareFunc


func setupRoutes() {
	r.Handle("/login/student", v1.StudentLogin()).Methods("POST")
	r.Handle("/login/teacher", v1.TeacherLogin()).Methods("POST")
	r.Handle("/Dashboard/teacher", mw.Auth(v1.Tdashboard())).Methods("GET")
	r.Handle("/Dashboard/student", mw.Auth(v1.Sdashboard())).Methods("GET")
}

func main() {
	router = mux.NewRouter()
	r = router.PathPrefix("/api").Subrouter()
	mwFunc = mw.InitializeCorsMw(r)
	setupRoutes()
	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
