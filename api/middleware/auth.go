package middleware

import (
	"fmt"
	"net/http"

	"github.com/vashish1/OnlineClassPortal/api/utility"
	"github.com/vashish1/OnlineClassPortal/pkg/database/student"
	"github.com/vashish1/OnlineClassPortal/pkg/database/teacher"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("executing middleware for authorization")
		str := r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")

		m, ok := utility.VerifyJwt(str)
		if !ok {
			w.WriteHeader(400)
			w.Write([]byte(`{"error": "Authentication Failed"}`))
			return
		}
		user := m["type"].(float64)
		if user == 0 {
			ok = student.IsAvailable(m["uid"].(string))
		} else {
			ok = teacher.IsAvailable(m["uid"].(string))
		}
		if ok {
			next.ServeHTTP(w, r)
		}
		w.WriteHeader(400)
		w.Write([]byte(`{"error": "Authentication Failed"}`))
		return
	})
}
