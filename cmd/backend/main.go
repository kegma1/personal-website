package main

import (
	"context"
	"database/sql"
	"log"
	"personal_website/internal"
	"personal_website/internal/backend"
	"personal_website/repositories"

	tea "charm.land/bubbletea/v2"

	_ "modernc.org/sqlite"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("sqlite", internal.Get_blog_database_path())
	if err != nil {
		log.Fatal(err)
	}

	q := repositories.New(db)

	app := backend.App{
		Ctx: ctx,
		Q: q,
	}

	p := tea.NewProgram(backend.InitialModel(&app))
	if _, err := p.Run(); err != nil {
		log.Fatalf("Error: %v\n", err)
	}
}

