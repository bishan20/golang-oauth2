package handlers

import (
	"context"
	"fmt"
	"golang-oauth2/utils"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:9091/callback",
		ClientID:     "YOUR_CLIENT_ID",
		ClientSecret: "YOUR_CLIENT_SECRET",
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
)

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {

	// generating oauth state and cookie
	oauthState := utils.GenerateOauthStateAndCookie(w)

	// oauth state is passed to AuthCodeURL, which is a token to protect user from CSRF attacks.
	url := googleOauthConfig.AuthCodeURL(oauthState)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallBackFromGoogle(w http.ResponseWriter, r *http.Request) {

	// get oauth state from cookie
	oauthState, _ := r.Cookie("oauthstate")

	if r.FormValue("state") != oauthState.Value {
		fmt.Println("state is not valid")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	token, err := googleOauthConfig.Exchange(context.Background(), r.FormValue("code"))
	if err != nil {
		fmt.Printf("could not get token: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		fmt.Printf("could not create get request: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("could not parse response: %s\n", err.Error())
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	fmt.Fprintf(w, "ParsedResponse: "+string(response)+"\n")
}
