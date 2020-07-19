package middleware

import "github.com/vashish1/OnlineClassPortal/api/utility"

func AuthMw(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("executing middleware for authorization")
		str := r.Header.Get("Authorization")
		w.Header().Set("Content-Type", "application/json")

		ok:=utility.VerifyJwt(str)
		if !ok{
			w.WriteHeader(resp.Status_code)
			b, _ := json.Marshal(resp)
			w.Write(b)
			return
		}
		fmt.Println("Verified ID token:", token)
		w.WriteHeader(http.StatusOK)
		next.ServeHTTP(w, r)
	}, BadExpr,

		BadExpr,
	)
}