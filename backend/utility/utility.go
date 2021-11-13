package utility

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/storage"
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

func UploadFile(object string, file multipart.File) error {
	b, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	tmpFile, err := os.OpenFile(object, os.O_RDWR|os.O_CREATE, 0755)
	_, err = tmpFile.Write(b)
	defer tmpFile.Close()

	bucket := "batbuck"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, tmpFile); err != nil {
		return err
	}
	if err := wc.Close(); err != nil {
		return err
	}
	return nil
}

func GenerateUUID() int {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 100000
	id := rand.Intn(max-min+1) + min
	return id
}
