package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
)

var conf *oauth2.Config
var tok *oauth2.Token
var err error

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func login(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		fmt.Println("Unable to read client secret file: %v", err)
	}
	conf, err = google.ConfigFromJSON(b, calendar.CalendarReadonlyScope,calendar.CalendarScope,calendar.CalendarEventsScope)
	if err != nil {
		fmt.Println("Unable to parse client secret file to config: %v", err)
	}
	fmt.Println(conf)
	authURL := conf.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Println("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)
		w.Write([]byte("<html><title>Golang Google</title> <body> <a href='" + authURL + "'><button>Login with Google!</button> </a> </body></html>"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/login", login).Methods("GET")
	r.HandleFunc("/auth", authhandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8000", nil)

}

func authhandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("okay")
	q := r.URL.Query()

	authcode := q.Get("code")
	tok, err := conf.Exchange(oauth2.NoContext, authcode)
	if err != nil {
		json.NewEncoder(w).Encode(tok)
		return
	}
	client := getClient(conf)
	srv, err := calendar.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	event := &calendar.Event{
		Summary:     "Google I/O 2021",
		Location:    "800 Howard St., San Francisco, CA 94103",
		Description: "A chance to hear more about Google's developer products.",
		Start: &calendar.EventDateTime{
			DateTime: "2021-05-28T09:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		End: &calendar.EventDateTime{
			DateTime: "2021-05-28T17:00:00-07:00",
			TimeZone: "America/Los_Angeles",
		},
		Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
		Attendees: []*calendar.EventAttendee{
			&calendar.EventAttendee {Email: "vashishtiv@gmail.com"},
			&calendar.EventAttendee{Email: "pooonamgupta8@gmail.com"},
		},
	}

	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)
	fmt.Print("done")
}
