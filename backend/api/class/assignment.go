package class

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/vashish1/OCLS/backend/database"
	"github.com/vashish1/OCLS/backend/models"
	"github.com/vashish1/OCLS/backend/utility"
)

// Content-Type: application/pdf
func GiveAssignment(w http.ResponseWriter, r *http.Request) {
	email := r.Context().Value("email")
	var res models.Response
	var code int

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

	ok, id := database.InsertAssignment(desc, t, h.Filename)
	if ok {
		if err := database.UpdateTeacher(email.(string), "assignment", id); err == nil {
			res = models.Response{
				Success: true,
				Message: "Assignment Added",
			}
			code = http.StatusAccepted
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
	email := r.Context().Value("email").(string)
	var res models.Response
	var code int

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

	ok := database.InsertSubmission(id, email, h.Filename)
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
