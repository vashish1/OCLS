package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/OnlineClassPortal/pkg/models"
	"github.com/vashish1/OnlineClassPortal/api/utility"
)

func StudentLogin() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var login models.Login
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &login)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "body not parsed"}`))
			return
		}
		var t models.LoginResponse

	   token,err:= utility.GenerateJwtForStudent(login.Email,login.Password)
	   if err!=nil{
		   t.Success=false
		   t.Token= ""
		w.WriteHeader(500)
		w.Write([]byte(t))
		return
	   }
	   t.Success=true
	   t.Token=token
	   json.NewEncoder(w).Encode(t)
	})
}

func CheckLogin() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		ok:=utility.VerifyJwtForStudent(tokenString)
		json.NewEncoder(w).Encode(ok)
	})
}