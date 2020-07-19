package v1

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/OnlineClassPortal/api/utility"
	"github.com/vashish1/OnlineClassPortal/pkg/models"
)

func TeacherLogin() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var login models.Login
		body, _ := ioutil.ReadAll(r.Body)
		err := json.Unmarshal(body, &login)
		fmt.Println("login", login)
		if err != nil {

			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error": "body not parsed"}`))
			return
		}
		var t models.LoginResponse

		token, err := utility.GenerateJwtForTeacher(login.Email, login.Password)
		if err != nil {
			fmt.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			t.Success = false
			t.Token = ""
			b, er := json.Marshal(t)
			fmt.Println(er)
			w.Write(b)
			return

		}
		t.Success = true
		t.Token = token
		json.NewEncoder(w).Encode(t)
	})
}
