package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/OCLS/database"
	"github.com/vashish1/OCLS/models"
	"github.com/vashish1/OCLS/utility"
)

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var input map[string]interface{}
	email := r.Context().Value("email").(string)
	user_type := (int)(r.Context().Value("type").(float64))
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
	var ok bool
	var output map[string]interface{}
	if user_type == models.Type_Student {
		ok, output = database.UpdateStudentDetails(email, input)
	} else {
		ok, output = database.UpdateTeacherDetails(email, input)
	}

	if ok {
		res = models.Response{
			Success: true,
			Message: "user updated successfully",
			Data:    output,
		}
		code = http.StatusOK
	} else {
		res = models.Response{
			Success: false,
			Error:   "error while updating user",
		}
		code = http.StatusBadRequest
	}

	utility.SendResponse(w, res, code)
	return
}
