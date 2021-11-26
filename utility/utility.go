package utility

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/vashish1/OCLS/models"
)

func SendResponse(w http.ResponseWriter, data models.Response, code int) {
	w.WriteHeader(code)
	b, _ := json.Marshal(data)
	w.Write(b)
	return
}

func SHA256ofstring(p string) string {
	h := sha1.New()
	h.Write([]byte(p))
	hash := hex.EncodeToString(h.Sum(nil))
	return hash
}

func UploadFile(object string, file multipart.File) error {

	bucket := "batbuck"
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		fmt.Println("3", err)
		return err
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	// Upload an object with storage.Writer.
	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
	if _, err = io.Copy(wc, file); err != nil {
		fmt.Println("1", err)
		return err
	}
	if err := wc.Close(); err != nil {
		fmt.Println("2", err)
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

func CreateSheet(data []models.Submission, sheet_type int) *excelize.File {
	f := excelize.NewFile()
	if sheet_type == models.Type_Written {
		f.SetCellValue("Sheet1", "A1", "File Name")
		f.SetCellValue("Sheet1", "B1", "Name")
		f.SetCellValue("Sheet1", "C1", "Email")
		f.SetCellValue("Sheet1", "D1", "Timestamp")
		f.SetCellValue("Sheet1", "E1", "Score")
		for i := 0; i < len(data); i++ {
			f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), data[i].FileName)
			f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), data[i].Name)
			f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), data[i].Email)
			f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), data[i].Timestamp)
			f.SetCellValue("Sheet1", "E"+strconv.Itoa(i+2), data[i].Score)
		}
		return f
	}
	f.SetCellValue("Sheet1", "A1", "Name")
	f.SetCellValue("Sheet1", "B1", "Email")
	f.SetCellValue("Sheet1", "C1", "Timestamp")
	f.SetCellValue("Sheet1", "D1", "Score")
	for i := 0; i < len(data); i++ {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(i+2), data[i].Name)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(i+2), data[i].Email)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(i+2), data[i].Timestamp)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(i+2), data[i].Score)
	}
	return f
}
