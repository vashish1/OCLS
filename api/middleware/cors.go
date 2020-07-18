package middleware

import "github.com/gorilla/mux"

func InitializeCorsMw(r *mux.Router) mux.MiddlewareFunc {
	MwFunc := mux.CORSMethodMiddleware(r)
	return MwFunc
}
