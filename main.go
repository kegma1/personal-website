package main

import (
	"log"
	"net/http"
	"personal_website/handlers"
)

func main() {
	mux := http.NewServeMux()
	
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/{path...}", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("GET /", handlers.TestStyle)
	
	log.Println("Server staring on :6969")
	log.Fatal(http.ListenAndServe(":6969", mux))
}
