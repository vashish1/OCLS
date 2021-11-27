package class

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/OCLS/api/Notification"
	"github.com/vashish1/OCLS/database"
	"github.com/vashish1/OCLS/models"
	"github.com/vashish1/OCLS/utility"
)

func CreateMCQ(w http.ResponseWriter, r *http.Request) {
	utility.EnableCors(&w)
	email, name, res, code := get(r)
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
		res.Error = err.Error()
		res.Success = false
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	ok := database.InsertMcq(input.Mcq, input.Date, input.Code, input.Desc, email, name)
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

func get(r *http.Request) (string, string, models.Response, int) {
	
	email := r.Context().Value("email").(string)
	name := r.Context().Value("name")
	var user_name string
	if name == nil {
		user_name = ""
	} else {
		user_name = name.(string)
	}
	var res models.Response
	var code int
	fmt.Println("here",email,name)
	return email, user_name, res, code
}

func SubmitMcq(w http.ResponseWriter, r *http.Request) {
	email, name, res, code := get(r)
	w.Header().Set("Content-Type", "application/json")
	var input struct {
		ID  int
		Ans []string
	}
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		res.Error = err.Error()
		res.Success = false
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	ok := database.InsertMcqSubmission(input.ID, input.Ans, email, name)
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
