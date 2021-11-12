package utility

import (
	"encoding/json"
	"net/http"

	"github.com/vashish1/OCLS/backend/models"
)

func SendResponse(w http.ResponseWriter, data models.Response, code int) {
	b, _ := json.Marshal(data)
	w.Write(b)
	w.WriteHeader(code)
	return
}

