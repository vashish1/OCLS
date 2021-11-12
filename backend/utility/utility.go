package utility

import (
	"crypto/sha1"
	"encoding/hex"
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

func SHA256ofstring(p string) string {
	h := sha1.New()
	h.Write([]byte(p))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}
