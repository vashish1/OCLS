package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/vashish1/OCLS/backend/api/auth"
	"github.com/vashish1/OCLS/backend/api/class"
	"github.com/vashish1/OCLS/backend/middleware"
)

func main() {
	fmt.Println("ok running")
	r := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"})
	methods := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	r.HandleFunc("/login", auth.Login).Methods("GET")
	r.HandleFunc("/signup", auth.Signup).Methods("POST")
	r.Handle("/class", middleware.Mdw.ThenFunc(class.CreateClass)).Methods("POST")
	http.Handle("/", handlers.CORS(headers, methods, origins)(r))
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	http.ListenAndServe(":"+port, nil)

}
