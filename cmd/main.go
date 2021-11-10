// Sample storage-quickstart creates a Google Cloud Storage bucket.
package main

import (
	"fmt"
	"io/ioutil"

	// "os"
	"time"

	"cloud.google.com/go/storage"
)

func main() {
	pkey, err := ioutil.ReadFile("key.pem")
	if err != nil {
		fmt.Println(err)
	}
	url, err := storage.SignedURL("batbuck", "ob1", &storage.SignedURLOptions{
		GoogleAccessID: "service-account-for-engage@engage-331409.iam.gserviceaccount.com",
		PrivateKey:     pkey,
		Method:         "GET",
		Expires:        time.Now().Add(48 * time.Hour),
	})
	if err != nil {
                fmt.Println(err)
	}
	fmt.Println(url)

	// ctx := context.Background()

	// // Sets your Google Cloud Platform project ID.
	// // projectID := "engage-331409"
	// jsonPath := "../cmd/engage-service-account-creds.json"
	// fmt.Println(jsonPath)
	// // Creates a client.
	// client, err := storage.NewClient(ctx, option.WithCredentialsFile(jsonPath))
	// if err != nil {
	// 	log.Fatalf("Failed to create client: %v", err)
	// }
	// defer client.Close()

	// // Open local file.
	// f, err := os.Open("notes.txt")
	// if err != nil {
	// 	fmt.Errorf("os.Open: %v", err)
	// }
	// defer f.Close()

	// ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	// defer cancel()
	// object := "ob1"
	// // Upload an object with storage.Writer.
	// wc := client.Bucket("batbuck").Object(object).NewWriter(ctx)
	// if _, err = io.Copy(wc, f); err != nil {
	// 	fmt.Errorf("io.Copy: %v", err)
	// }
	// if err := wc.Close(); err != nil {
	// 	fmt.Errorf("Writer.Close: %v", err)
	// }
	// fmt.Println("Blob %v uploaded.\n", object)

}
