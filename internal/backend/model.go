package backend

import (
	"log"
	"personal_website/repositories"

	tea "charm.land/bubbletea/v2"
)

type Model struct {
	app      *App
	projects []repositories.Project
}

func (m Model) Init() tea.Cmd {
	return nil
}

func InitialModel(app *App) Model {
	initalProjects, err := app.Q.GetProjects(app.Ctx)
	if err != nil {
		log.Fatal(err)
	}
	return Model{
		app:      app,
		projects: initalProjects,
	}
}
