package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	db "github.com/vashish1/OCLS/backend/database"
	"github.com/vashish1/OCLS/backend/middleware"
	"github.com/vashish1/OCLS/backend/models"
	"github.com/vashish1/OCLS/backend/utility"
)

type login struct {
	Email string `json:"email,omitempty"`
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
		w.Write([]byte(`{"error": "body not parsed"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Check if the user exissts with such credentials
	//donon ka check krna padega student aur teacher phir jwt uske acc create krna hai
	ok, user := db.UserExists(input.Email)
	if ok {
		tokenstring, err := middleware.GenerateAuthToken(input.Email, (int)(user["type"].(float64)))
		if err != nil {
			res = models.Response{
				Success: false,
				Message: "Incorrect Credentials, Try Again. ",
				Error:   err.Error(),
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
		Message: "no such user exist",
	}
	utility.SendResponse(w, res, code)
	return
}
