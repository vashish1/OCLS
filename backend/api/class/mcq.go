package class

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/OCLS/backend/api/Notification"
	"github.com/vashish1/OCLS/backend/database"
	"github.com/vashish1/OCLS/backend/models"
	"github.com/vashish1/OCLS/backend/utility"
)

func CreateMCQ(w http.ResponseWriter, r *http.Request) {
	email, res, code := get(r)
	w.Header().Set("Content-Type", "application/json")
	var input struct {
		Mcq  models.Mcq
		Code string
		Desc string
		Date string
	}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		w.Write([]byte(`{"error": "body not parsed"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ok := database.InsertMcq(input.Mcq, input.Date, input.Code, input.Desc, email.(string))
	if ok {
		res = models.Response{
			Success: true,
			Message: "Assignment Added",
		}
		code = http.StatusAccepted
		ok := Notification.SendEmail(input.Code, input.Date)
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

func get(r *http.Request) (interface{}, models.Response, int) {
	email := r.Context().Value("email")
	var res models.Response
	var code int
	return email, res, code
}

func SubmitMcq(w http.ResponseWriter, r *http.Request) {
	email, res, code := get(r)
	w.Header().Set("Content-Type", "application/json")
	var input struct {
		ID  int
		Ans []string
	}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		w.Write([]byte(`{"error": "body not parsed"}`))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	ok := database.InsertMcqSubmission(input.ID, input.Ans, email.(string))
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
