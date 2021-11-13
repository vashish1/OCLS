package class

import (
	"net/http"

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
	var desc=r.FormValue("description")
    ok,id :=database.InsertAssignment(desc,h.Filename)
	if ok{
		if err :=database.UpdateTeacher(email.(string),"assignment",id); err==nil{
                res=models.Response{
					Success: true,
					Message: "Assignment Added",
				}
				code=http.StatusAccepted
		}
	}else{
		res=models.Response{
			Success: false,
			Error: err.Error(),
		}
		code=http.StatusBadRequest
	}
	utility.SendResponse(w,res,code)
	return

}

func SubmitAssignment() {

}

func GetList() {

}
