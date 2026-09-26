package main

import (
	tea "charm.land/bubbletea/v2"
	"context"
	"database/sql"
	"log"
	"os"
	"personal_website/repositories"

	_ "modernc.org/sqlite"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite", get_blog_database_path())
	if err != nil {
		log.Fatal(err)
	}

	q := repositories.New(db)

}

func get_blog_database_path() string {
	if path := os.Getenv("BLOG_DATABASE_PATH"); path != "" {
		return path
	}
	return "./data/db.db"
}
