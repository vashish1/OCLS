package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/vashish1/OCLS/models"
	"github.com/vashish1/OCLS/utility"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type credential struct {
	Cid          string   `json:"client_id"`
	Csecret      string   `json:"client_secret"`
	Origins      []string `json:"javascript_origins"`
	Redirect     []string `json:"redirect_uris"`
	AuthProvider string   `json:"auth_provider_x509_cert_url"`
	ProjectId    string   `json:"project_id"`
	Auth_URI     string   `json:"auth_uri"`
	Token_URI    string   `json:"token_uri"`
}

var (
	cred              credential
	googleOauthConfig *oauth2.Config
	randomState       = "random"
)

func init() {

	f, err := ioutil.ReadFile("./api/auth/creds.json")
	if err != nil {
		fmt.Println("could not read the file:", err)
	}
	err = json.Unmarshal(f, &cred)
	// fmt.Print(cred.Redirect[0])
	googleOauthConfig = &oauth2.Config{

		RedirectURL:  cred.Redirect[1],
		ClientID:     cred.Cid,
		ClientSecret: cred.Csecret,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
			// "https://www.googleapis.com/auth/calendar",
			// "https://www.googleapis.com/auth/calendar.events",
		},
		Endpoint: google.Endpoint,
	}
}

func GoogleSignupHandler(w http.ResponseWriter, r *http.Request) {
	// url := googleOauthConfig.AuthCodeURL(randomState)
	URL, err := url.Parse(googleOauthConfig.Endpoint.AuthURL)
	if err != nil {
		fmt.Println("Parse: " + err.Error())
	}
	parameters := url.Values{}
	parameters.Add("client_id", googleOauthConfig.ClientID)
	parameters.Add("scope", strings.Join(googleOauthConfig.Scopes, " "))
	parameters.Add("redirect_uri", googleOauthConfig.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", randomState)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

//GoogleCallbackHandler func
func GoogleCallbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Print("callback recieved")
	// var res utility.Result

	if r.FormValue("state") != randomState {
		w.Write([]byte(`state invalid`))
		return
	}

	token, err := googleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	if err != nil {
		w.Write([]byte(`token invalid`))
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		w.Write([]byte(`state response`))
		return
	}

	defer resp.Body.Close()

	var user struct {
		ID      string `json:"id"`
		Email   string `json:"email"`
		Name    string `json:"name"`
		Fname   string `json:"given_name"`
		Sname   string `json:"family_name"`
		Picture string `json:"picture"`
	}

	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		w.Write([]byte(`decoding invalid`))
		return
	}
	
	URL, err := url.Parse("https://thawing-mountain-02190.herokuapp.com/")
	if err != nil {
		fmt.Println("Parse: " + err.Error())
	}
	parameters := url.Values{}
	parameters.Add("name", user.Name)
	parameters.Add("email", user.Email)
	URL.RawQuery = parameters.Encode()
	url := URL.String()
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}

func Welcome(w http.ResponseWriter,r *http.Request){
         q:=r.URL.Query()
		 var output struct {
			Name  string
			Email string
		}
		output.Name = q.Get("name")
		output.Email = q.Get("email")
		res:=models.Response{
			Success: true,
			Data: output,
		}
		utility.SendResponse(w,res,200)
	return
}