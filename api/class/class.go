package class

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/OCLS/database"
	"github.com/vashish1/OCLS/models"
	"github.com/vashish1/OCLS/utility"
)

func CreateClass(w http.ResponseWriter, r *http.Request) {
	utility.EnableCors(&w)
	user_type := r.Context().Value("type")
	email, name, res, code := get(r)
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
		res = models.Response{
			Success: false,
			Message: "Error while Reading Request body",
		}
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	input.TeacherEmail = email
	input.TeacherName = name
	input.StudentList=[]models.List{}
	ok, class_code := database.InsertClass(input, email)
	fmt.Println(ok, class_code)
	if ok {
		res = models.Response{
			Success: true,
			Message: "Class created Successfully",
			Data:    class_code,
		}
		code = http.StatusAccepted
		utility.SendResponse(w, res, code)
		return
	}
	res = models.Response{
		Success: false,
		Message: "Error while creating class,try using another name",
	}
	code = http.StatusBadRequest

	utility.SendResponse(w, res, code)
	return
}

func JoinClass(w http.ResponseWriter, r *http.Request) {
	utility.EnableCors(&w)
	user_type := r.Context().Value("type")
	email, name, res, code := get(r)
	if (int)(user_type.(float64)) != models.Type_Student {
		res = models.Response{
			Success: false,
			Error:   "unauthorized user for request",
		}
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	var input struct {
		Class_Code string `json:"class_code,required"`
	}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		res = models.Response{
			Success: false,
			Error:   "Error while Reading Request body",
		}
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}

	ok := database.UpdataClassData(input.Class_Code, email, name)
	if ok {
		res = models.Response{
			Success: true,
			Message: "Class Joined Successfully",
		}
		code = http.StatusAccepted
	} else {
		res = models.Response{
			Success: false,
			Error:   "Error while joining the class",
		}
		code = http.StatusForbidden
	}
	utility.SendResponse(w, res, code)
	return

}

func GetClass(w http.ResponseWriter, r *http.Request) {
	utility.EnableCors(&w)
    email, _, res, code := get(r)
	user_type := (int)(r.Context().Value("type").(float64))
	ok, data := database.GetAllClass(email,user_type)
	if ok {
		res := models.Response{
			Success: true,
			Message: "class data fetch successful",
			Data:    data,
		}
		code=http.StatusOK
		utility.SendResponse(w, res, code)
		return
	}
	res = models.Response{
		Success: false,
		Error:   "error while fetching data",
	}
	code=http.StatusInternalServerError
	utility.SendResponse(w, res, code)
	return
}
