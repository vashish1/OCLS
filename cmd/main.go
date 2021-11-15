package main

// func middleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("entering middleware")
// 		next.ServeHTTP(w, r)
// 	})
// }

// func index(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseMultipartForm(32 << 20) // maxMemory 32MB
// 	if err != nil {
// 	   fmt.Println(err)
// 		w.WriteHeader(400)
// 		return
// 	}
// 	file, _, err := r.FormFile("file")
// 	if err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(400)
// 		return
// 	}

// 	bucket := "batbuck"
// 	object := "ob2"
// 	ctx := context.Background()
// 	client, err := storage.NewClient(ctx)
// 	if err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(400)
// 		return
// 	}
// 	defer client.Close()

// 	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
// 	defer cancel()

// 	// Upload an object with storage.Writer.
// 	wc := client.Bucket(bucket).Object(object).NewWriter(ctx)
// 	if _, err = io.Copy(wc, file); err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(400)
// 		return
// 	}
// 	if err := wc.Close(); err != nil {
// 		fmt.Println(err)
// 		w.WriteHeader(400)
// 		return
// 	}
// 	return
// }

// func main() {

// 	ok := Notification.SendWelcomeEmail("vashishtiv@gmail.com", "Yashi Gupta", "")
// 	fmt.Println(ok)
// }
