package main

import (
	"log"
	"net/http"
	"personal_website/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", handlers.TestStyle)
	
	log.Println("Server staring on :6969")
	log.Fatal(http.ListenAndServe(":6969", mux))
}
