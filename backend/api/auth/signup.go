package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/vashish1/OCLS/backend/models"
	db "github.com/vashish1/OCLS/backend/database"
)

func signup(w http.ResponseWriter, r *http.Request) {
	var input models.Student
	w.Header().Set("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(body, &input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "body not parsed"}`))
		return
	}
	ok, err := db.Insertintodb(input)
	var res models.Response
	var code int
	if ok {
		w.WriteHeader(http.StatusOK)
		res =models.Response{
				Message: "Signup SuccessFul",
				Success: true,
				Error: "",
				}
			code=http.StatusAccepted
	} else {
			res = models.Response{
				Error: err.Error(),
			}
			code=http.StatusBadRequest
	}
	b, _ := json.Marshal(res)
	w.Write(b)
	w.WriteHeader(code)	
	 return
}
