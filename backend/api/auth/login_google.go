package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	db "github.com/vashish1/OCLS/backend/database"
	"github.com/vashish1/OCLS/backend/middleware"
	"github.com/vashish1/OCLS/backend/models"
	"github.com/vashish1/OCLS/backend/utility"
)

type login_google struct {
	Name string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
}

func LoginGoogle(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var input login_google
	var res models.Response
	var code int
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		res = models.Response{
			Success: false,
			Error: err.Error(),
		}
		code = http.StatusBadRequest
     	utility.SendResponse(w, res, code)
		return
	}
	ok,user := db.CheckEmail(input.Email)
	if ok {
		tokenstring, err := middleware.GenerateAuthToken(input.Email,user["name"].(string), (int)(user["type"].(float64)))
		if err != nil {
			fmt.Println(err.Error())
			res = models.Response{
				Success: false,
				Error: "Incorrect Credentials, Try Again. ",
			}
			code = http.StatusBadRequest
		} else {
			//Send a Successfull Response
			res = models.Response{
				Message: "Log In successful",
				Success: true,
				Data:    tokenstring,
			}
			code = http.StatusAccepted
		}
		utility.SendResponse(w, res, code)
		return
	}

	res = models.Response{
		Success: false,
		Error: "no such user exist",
	}
	code=http.StatusBadRequest
	utility.SendResponse(w, res, code)
	return
}