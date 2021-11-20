package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/vashish1/OCLS/models"
	"github.com/vashish1/OCLS/utility"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("entering middleware")
		var res models.Response
		var code int
		w.Header().Set("Content-Type", "application/json")
		authStr := r.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authStr, "Bearer ")
		if tokenString != "" {
			ok, details := VerifyAuthToken(tokenString)
			if ok {
				//Todo add details to context
				con := context.WithValue(r.Context(), "type", details["type"])
				newCon := context.WithValue(con, "email", details["email"])
				next.ServeHTTP(w, r.WithContext(newCon))
				return
			}
			res = models.Response{
				Success: false,
				Message: "Unauthorized",
			}
			code = http.StatusForbidden
		} else {
			res = models.Response{
				Success: false,
				Message: "Authentication string required",
			}
			code = http.StatusBadRequest
		}
		utility.SendResponse(w, res, code)
		return
	})
}
