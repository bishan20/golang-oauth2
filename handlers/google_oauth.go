package handlers

import (
	"context"
	"fmt"
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
	oauthStateString = "thisisrandom"
)

func handleGoogleLogin(w http.ResponseWriter, r *http.Request) {

	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	fmt.Println(url)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func handleCallBackFromGoogle(w http.ResponseWriter, r *http.Request) {

	if r.FormValue("state") != oauthStateString {
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
