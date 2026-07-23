package main

import (
	"context"
	"database/sql"
	"log"
	"personal_website/repositories"

	_ "modernc.org/sqlite"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite", "./db.db")
	if err != nil {
		log.Fatal(err)
	}

	repo := repositories.New(db)

	post, err := repo.CreatePost(ctx, repositories.CreatePostParams{
		Title:    "test",
		Subtitle: "subtest",
		Content:  "content test",
		Tags:     "test, test2",
		Language: "Norsk",
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println(post)
}
