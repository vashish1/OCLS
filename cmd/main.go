package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func middleware(next http.Handler)(http.Handler){
 return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
	 fmt.Println("entering middleware")
	next.ServeHTTP(w,r)
 })
}

func index(w http.ResponseWriter, r *http.Request) {
fmt.Println("reached the function")
}

func main() {
	r := mux.NewRouter()
	mdw := alice.New(middleware)
	r.Handle("/", mdw.ThenFunc(index))
	http.ListenAndServe(":8000", r)
}
