package utils

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"
)

/*
Generating oauth state string and sending it to http response header.
*/
func GenerateOauthStateAndCookie(w http.ResponseWriter) string {

	cookieExpirationTime := time.Now().Add(2 * time.Minute)

	b := make([]byte, 16)

	rand.Read(b)
	state := base64.StdEncoding.EncodeToString(b)

	cookie := http.Cookie{
		Name:     "oauthstate",
		Value:    state,
		Expires:  cookieExpirationTime,
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	return state
}
