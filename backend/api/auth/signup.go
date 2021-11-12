package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	db "github.com/vashish1/OCLS/backend/database"
	"github.com/vashish1/OCLS/backend/models"
	"github.com/vashish1/OCLS/backend/utility"
)

//email verification bhi add krni hai student aur teacher ke liye
//pass encrypt bhi krna hai
func Signup(w http.ResponseWriter, r *http.Request) {
	var input map[string]interface{}
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	ok, err := db.Insertintodb(input)
	var res models.Response
	var code int
	if ok {
		w.WriteHeader(http.StatusOK)
		res = models.Response{
			Message: "Signup SuccessFul",
			Success: true,
			Error:   "",
		}
		code = http.StatusAccepted
	} else {
		res = models.Response{
			Error: err.Error(),
		}
		code = http.StatusBadRequest
	}
	utility.SendResponse(w, res, code)
	return
}
