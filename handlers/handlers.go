package handlers

import (
	"context"
	"log"
	"net/http"
	"personal_website/components"
	"personal_website/repositories"
	"strconv"
)

type Application struct {
	Ctx context.Context
	Q *repositories.Queries
}

func (app *Application) Desktop(w http.ResponseWriter, r *http.Request) {
	component := components.Desktop()
	component.Render(r.Context(), w)	
}

func (app *Application) Posts(w http.ResponseWriter, r *http.Request) {
	page, err := strconv.Atoi(r.PathValue("page"))
	if err != nil {
		log.Fatal(err)
	}

	posts, err := app.Q.GetPagedPosts(app.Ctx, repositories.GetPagedPostsParams{ Limit: 5, Offset: int64(page * 5) })
	if err != nil {
		log.Fatal(err)
	}

	log.Println(page)

	component := components.PostPage(page + 1, posts)
	component.Render(r.Context(), w)
}
