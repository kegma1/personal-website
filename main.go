package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"personal_website/handlers"
	"personal_website/repositories"

	_ "modernc.org/sqlite"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite", "./db.db")
	if err != nil {
		log.Fatal(err)		
	}

	q := repositories.New(db)

	mux := http.NewServeMux()

	a := &handlers.Application {
		Ctx: ctx,
		Q: q,
	}
	
	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("GET /static/{path...}", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("GET /", a.Desktop)
	mux.HandleFunc("GET /posts/{page}", a.Posts)
	
	log.Println("Server staring on :6969")
	log.Fatal(http.ListenAndServe(":6969", mux))
}
