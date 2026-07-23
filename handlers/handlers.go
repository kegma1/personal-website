package handlers

import (
	"context"
	"log"
	"net/http"
	"personal_website/components"
	"personal_website/repositories"
)

type Application struct {
	Ctx context.Context
	Q *repositories.Queries
}

func (app *Application) Desktop(w http.ResponseWriter, r *http.Request) {
	posts, err := app.Q.GetPosts(app.Ctx) 
	if err != nil {
		log.Fatal(err)
	}

	log.Println(posts)
	component := components.Desktop()
	component.Render(r.Context(), w)	
}
