package main

import (
	"golang-oauth2/handlers"
	"log"
	"net/http"
)

func main() {

	server := &http.Server{
		Addr:    ":9091",
		Handler: handlers.New(),
	}

	log.Printf("HTTP Server listening at %s", server.Addr)

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		log.Printf("%v", err)
	} else {
		log.Println("Server closed!")
	}
}
