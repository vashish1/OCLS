package class

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/OCLS/backend/database"
	"github.com/vashish1/OCLS/backend/models"
	"github.com/vashish1/OCLS/backend/utility"
)

func CreateClass(w http.ResponseWriter, r *http.Request) {
	user_type := r.Context().Value("type")
	email := r.Context().Value("email")
	var res models.Response
	var code int
	if (int)(user_type.(float64)) != models.Type_Teacher {
		res = models.Response{
			Success: false,
			Message: "unauthorized user for request",
		}
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var input models.Class
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		w.Write([]byte(`{"error": "body not parsed"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	input.TeacherEmail = email.(string)
	ok, class_code := database.InsertClass(input)
	fmt.Println(ok, class_code)
	if ok {
		err := database.UpdateTeacher(email.(string), class_code)
		fmt.Println(err)
		if err == nil {
			res = models.Response{
				Success: true,
				Message: "Class created Successfully",
				Data:    class_code,
			}
			code = http.StatusAccepted
			utility.SendResponse(w, res, code)
			return
		}
	}
	res = models.Response{
		Success: false,
		Message: "Error while creating class",
	}
	code = http.StatusBadRequest

	utility.SendResponse(w, res, code)
	return
}

func JoinClass(w http.ResponseWriter, r *http.Request) {
	user_type := r.Context().Value("type")
	email := r.Context().Value("email")
	var res models.Response
	var code int
	if (int)(user_type.(float64)) != models.Type_Teacher {
		res = models.Response{
			Success: false,
			Message: "unauthorized user for request",
		}
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var input struct {
		Class_Code string
	}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		w.Write([]byte(`{"error": "body not parsed"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	ok := database.UpdataClassData(input.Class_Code, email.(string))
	if ok {
		res = models.Response{
			Success: true,
			Message: "Class Joined Successfully",
		}
		code = http.StatusAccepted
	} else {
		res = models.Response{
			Success: false,
			Message: "Error while joining the class",
		}
		code = http.StatusForbidden
	}
	utility.SendResponse(w, res, code)
	return

}
