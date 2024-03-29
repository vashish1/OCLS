package class

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/vashish1/OCLS/api/Notification"
	"github.com/vashish1/OCLS/database"
	"github.com/vashish1/OCLS/models"
	"github.com/vashish1/OCLS/utility"
)

// Content-Type: application/pdf
func CreateAssignment(w http.ResponseWriter, r *http.Request) {
	utility.EnableCors(&w)
	email, name, res, code := get(r)
	fmt.Println(name)
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
	fmt.Println(t)
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
		utility.SendResponse(w, res, code)
		return
	}

	ok := database.InsertAssignment(desc, t, h.Filename, class_code, email, name)
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
	fmt.Println("here")
	utility.SendResponse(w, res, code)
	return
}

func SubmitAssignment(w http.ResponseWriter, r *http.Request) {
	utility.EnableCors(&w)
	email, name, res, code := get(r)
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

	ok := database.InsertSubmission(id, email, name, h.Filename)
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
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	id_value, _ := strconv.Atoi(id)
	fmt.Println(id)
	var res models.Response
	var code int
	err, data, _ := database.GetSubmissions(id_value)
	if err != nil {
		res.Error = err.Error()
		res.Success = false
		code = http.StatusBadRequest
	} else {
		res.Data = data
		res.Message = "successful"
		res.Success = true
		code = http.StatusOK
	}
	utility.SendResponse(w, res, code)
	return
}

func DownloadSubmission(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	id_value, _ := strconv.Atoi(id)
	_, data, sheet_type := database.GetSubmissions(id_value)
	file := utility.CreateSheet(data, sheet_type)

	// Set the headers necessary to get browsers to interpret the downloadable file
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\"submission.xlsx\"")
	w.Header().Set("File-Name", "submission.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	err := file.Write(w)
	fmt.Println(err)
}

func GetAssignment(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Class string `json:"class,omitempty"`
	}
	var res models.Response
	var code int
	w.Header().Set("Content-Type", "application/json")
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
	ok, data := database.GetAllAssignment(input.Class,)
	if ok {
		temp:=Notification.MarshalAssignment(data)
		res = models.Response{
			Success: true,
			Message: "class data fetch successful",
			Data:    temp,
		}
		utility.SendResponse(w, res, http.StatusOK)
		return
	}
	res = models.Response{
		Success: false,
		Error:   "error while fetching data",
	}
	utility.SendResponse(w, res, http.StatusInternalServerError)
	return
}

func GetAssignmentStudent(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Class string `json:"class,omitempty"`
	}
	var res models.Response
	var code int
	w.Header().Set("Content-Type", "application/json")
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
	ok, data := database.GetAllAssignment(input.Class,)
	if ok {
		temp:=Notification.MarshalAssignment(data)
        var result []models.Assignment
        for _,x:=range temp{
			if (time.Now()).Before(x.Date){
				result = append(result, x)
			}
		}
		res = models.Response{
			Success: true,
			Message: "class data fetch successful",
			Data:    result,
		}
		utility.SendResponse(w, res, http.StatusOK)
		return
	}
	res = models.Response{
		Success: false,
		Error:   "error while fetching data",
	}
	utility.SendResponse(w, res, http.StatusInternalServerError)
	return
}