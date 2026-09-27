package backend

import (
	"log"
	"personal_website/repositories"

	tea "charm.land/bubbletea/v2"
)

type MainModel struct {
	app      *App
	projects []repositories.Project
	createProject CreateProjectModel
}

func (m MainModel) Init() tea.Cmd {
	return nil
}

func InitialModel(app *App) MainModel {
	initalProjects, err := app.Q.GetProjects(app.Ctx)
	if err != nil {
		log.Fatal(err)
	}
	return MainModel{
		app:      app,
		projects: initalProjects,
		createProject: initialCreateProjectModel(),
	}
}
