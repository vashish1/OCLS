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
	ok, class_code := database.InsertClass(input, email.(string))
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

func CreateAnnouncement(w http.ResponseWriter, r *http.Request) {
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
	var input models.Announcement
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		w.Write([]byte(`{"error": "body not parsed"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ok := database.InsertAnnouncement(input, email.(string))
	if ok {
		res = models.Response{
			Success: true,
			Message: "Announcement added Successfully",
		}
		code = http.StatusAccepted
	} else {
		res = models.Response{
			Success: false,
			Error:   "Error while adding announcement",
		}
		code = http.StatusAccepted
	}

}

func GetClass(w http.ResponseWriter, r *http.Request) {

	ok, data := database.GetAllClass()
	if ok {
		res := models.Response{
			Success: true,
			Message: "dlass data fetch successful",
			Data:    data,
		}
		utility.SendResponse(w, res, http.StatusOK)
		return
	}
	res := models.Response{
		Success: false,
		Error:   "error while fetching data",
	}
	utility.SendResponse(w, res, http.StatusInternalServerError)
	return
}

func GetAnnouncement(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Class string `json:"class,omitempty"`
	}
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		w.Write([]byte(`{"error": "body not parsed"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ok, data := database.GetAllAnnouncement(input.Class)
	if ok {
		res := models.Response{
			Success: true,
			Message: "dlass data fetch successful",
			Data:    data,
		}
		utility.SendResponse(w, res, http.StatusOK)
		return
	}
	res := models.Response{
		Success: false,
		Error:   "error while fetching data",
	}
	utility.SendResponse(w, res, http.StatusInternalServerError)
	return
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Class string `json:"class,omitempty"`
	}
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		w.Write([]byte(`{"error": "body not parsed"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ok, data := database.GetAllAssignment(input.Class)
	if ok {
		res := models.Response{
			Success: true,
			Message: "class data fetch successful",
			Data:    data,
		}
		utility.SendResponse(w, res, http.StatusOK)
		return
	}
	res := models.Response{
		Success: false,
		Error:   "error while fetching data",
	}
	utility.SendResponse(w, res, http.StatusInternalServerError)
	return
}
