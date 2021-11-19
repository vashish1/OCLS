package class

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vashish1/OCLS/backend/api/Notification"
	"github.com/vashish1/OCLS/backend/database"
	"github.com/vashish1/OCLS/backend/models"
	"github.com/vashish1/OCLS/backend/utility"
)

// Content-Type: application/pdf
func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	email, res, code := get(r)
	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		res.Error = err.Error()
		res.Success = false
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	var desc = r.FormValue("description")
	var t = r.FormValue("date")
	var class_code = r.FormValue("class_code")
	file, h, err := r.FormFile("file")
	if err != nil {
		res.Error = err.Error()
		res.Success = false
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	err = utility.UploadFile(h.Filename, file)
	if err != nil {
		res = models.Response{
			Error:   err.Error(),
			Success: false,
		}
		code = http.StatusInternalServerError
	}

	ok := database.InsertAssignment(desc, t, h.Filename, class_code, email.(string))
	if ok {
		res = models.Response{
			Success: true,
			Message: "Assignment Added",
		}
		code = http.StatusAccepted
		ok := Notification.SendEmail(class_code, t)
		if ok {
			fmt.Println("email sent")
		} else {
			fmt.Println("email not sent")
		}
	} else {
		res = models.Response{
			Success: false,
			Error:   err.Error(),
		}
		code = http.StatusBadRequest
	}
	utility.SendResponse(w, res, code)
	return
}

func SubmitAssignment(w http.ResponseWriter, r *http.Request) {
	email, res, code := get(r)
	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
	if err != nil {
		res.Error = err.Error()
		res.Success = false
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	id := r.FormValue("id")
	file, h, err := r.FormFile("file")
	if err != nil {
		res.Error = err.Error()
		res.Success = false
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	err = utility.UploadFile(h.Filename, file)
	if err != nil {
		res = models.Response{
			Error:   err.Error(),
			Success: false,
		}
		code = http.StatusInternalServerError
	}

	ok := database.InsertSubmission(id, email.(string), h.Filename)
	if ok {
		res = models.Response{
			Success: true,
			Message: "Assignment Submitted",
		}
		code = http.StatusAccepted
	} else {
		res = models.Response{
			Success: false,
			Error:   err.Error(),
		}
		code = http.StatusBadRequest
	}
	utility.SendResponse(w, res, code)
	return
}

func GetSubmissionList(w http.ResponseWriter, r *http.Request) {

	id := mux.Vars(r)["id"]
	id_value, _ := strconv.Atoi(id)
	var res models.Response
	var code int
	err, data := database.GetSubmissions(id_value)
	if err != nil {
		res.Error = err.Error()
		res.Success = false
		code = http.StatusBadRequest
	} else {
		res.Data = data
		res.Message = "successful"
		res.Success = true
		code = http.StatusBadRequest
	}
	utility.SendResponse(w, res, code)
	return
}
