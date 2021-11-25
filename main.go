package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vashish1/OCLS/api/auth"
	"github.com/vashish1/OCLS/api/class"
	"github.com/vashish1/OCLS/middleware"
)

func main() {
	fmt.Println("ok running")

	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*", "Access-Control-Allow-Origin"})
	r.HandleFunc("/login", auth.Login).Methods("POST")
	r.HandleFunc("/signup", auth.Signup).Methods("POST")
	r.HandleFunc("/signup/google", auth.GoogleSignupHandler).Methods("POST")
	r.HandleFunc("/login/google", auth.LoginGoogle).Methods("POST")
	r.HandleFunc("/callback", auth.GoogleCallbackHandler)
	r.Handle("/class", middleware.Mdw.ThenFunc(class.CreateClass)).Methods("POST")
	r.Handle("/class/add", middleware.Mdw.ThenFunc(class.CreateClass)).Methods("POST")
	r.Handle("/class/join", middleware.Mdw.ThenFunc(class.JoinClass)).Methods("POST")
	r.Handle("/class/get", middleware.Mdw.ThenFunc(class.GetClass)).Methods("GET")
	r.Handle("/class/announcement/add", middleware.Mdw.ThenFunc(class.CreateAnnouncement)).Methods("POST")
	r.Handle("/class/announcement/get", middleware.Mdw.ThenFunc(class.GetAnnouncement)).Methods("GET")
	r.Handle("/class/assignment/add", middleware.Mdw.ThenFunc(class.CreateAssignment)).Methods("POST")
	r.Handle("/class/assignment/get", middleware.Mdw.ThenFunc(class.GetAssignment)).Methods("GET")
	r.Handle("/class/assignment/sub", middleware.Mdw.ThenFunc(class.SubmitAssignment)).Methods("POST")
	r.Handle("/class/assignment/sub/{id}", middleware.Mdw.ThenFunc(class.GetSubmissionList)).Methods("GET")
	r.Handle("/submission/{id}", middleware.Mdw.ThenFunc(class.DownloadSubmission)).Methods("GET")
	r.Handle("/class/mcq/add", middleware.Mdw.ThenFunc(class.CreateMCQ)).Methods("POST")
	r.Handle("/class/mcq/sub", middleware.Mdw.ThenFunc(class.SubmitMcq)).Methods("POST")
	r.Handle("/user/update", middleware.Mdw.ThenFunc(auth.UpdateUser)).Methods("POST")

	http.Handle("/", handlers.CORS(headers, methods, origins)(r))
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	http.ListenAndServe(":"+port, nil)

}
