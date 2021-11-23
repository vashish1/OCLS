package class

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/OCLS/database"
	"github.com/vashish1/OCLS/models"
	"github.com/vashish1/OCLS/utility"
)

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

func CreateAnnouncement(w http.ResponseWriter, r *http.Request) {
	user_type := r.Context().Value("type")
	email, name, res, code := get(r)
	if (int)(user_type.(float64)) != models.Type_Teacher {
		res = models.Response{
			Success: false,
			Error:   "unauthorized user for request",
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
		res = models.Response{
			Success: false,
			Error:   "Error while Reading Request body",
		}
		code = http.StatusBadRequest
		utility.SendResponse(w, res, code)
		return
	}
	input.TeacherName = name
	ok := database.InsertAnnouncement(input, email)
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
    utility.SendResponse(w,res,code)
}
