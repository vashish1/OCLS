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

type login_google struct {
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
}

func LoginGoogle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input login_google
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
	ok, user := db.CheckEmail(input.Email)
	if ok {
		tokenstring, err := middleware.GenerateAuthToken(input.Email, user["name"].(string), (int)(user["type"].(float64)))
		if err != nil {
			fmt.Println(err.Error())
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
