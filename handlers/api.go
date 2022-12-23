package handlers

import "net/http"

func New() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("templates/")))

	mux.HandleFunc("/login", handleGoogleLogin)
	mux.HandleFunc("/callback", handleCallBackFromGoogle)

	return mux
}
