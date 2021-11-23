package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	db "github.com/vashish1/OCLS/database"
	"github.com/vashish1/OCLS/middleware"
	"github.com/vashish1/OCLS/models"
	"github.com/vashish1/OCLS/utility"
)

type login struct {
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

//login to Implement Login of user.
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var input login
	var res models.Response
	var code int
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		res = models.Response{
			Success: false,
			Error:   err.Error(),
		}
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	var name string

	ok, user := db.UserExists(input.Email, input.Password)
	if ok {
		if user["name"] == nil {
			name = ""
		} else {
			name = user["name"].(string)
		}
		tokenstring, err := middleware.GenerateAuthToken(input.Email, name, (int)(user["type"].(float64)))
		if err != nil {
			fmt.Println(err)
			res = models.Response{
				Success: false,
				Error:   "Incorrect Credentials, Try Again. ",
			}
			code = http.StatusBadRequest
		} else {
			var d struct{
				User map[string]interface{}
				Token string
			}
			d.User=user
			d.Token=tokenstring

			//Send a Successfull Response
			res = models.Response{
				Message: "Log In successful",
				Success: true,
				Data:    d,
			}
			code = http.StatusAccepted
		}
		utility.SendResponse(w, res, code)
		return
	}

	res = models.Response{
		Success: false,
		Error:   "no such user exist",
	}
	code = http.StatusBadRequest
	utility.SendResponse(w, res, code)
	return
}
