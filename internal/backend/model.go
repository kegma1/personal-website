package backend

import (
	"log"
	"personal_website/repositories"
	"strconv"

	"charm.land/bubbles/v2/table"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type MainModel struct {
	app           *App
	projects      []repositories.Project
	projectTable  table.Model
	createProject CreateProjectModel
	err           error
}

func (m MainModel) Init() tea.Cmd { return nil }

func (m *MainModel) updateRows() {
	rows := []table.Row{}
	for _, proj := range m.projects {
		rows = append(rows, table.Row{strconv.Itoa(int(proj.ID)), proj.Title, proj.Link, proj.IconPath})
	}
	m.projectTable.SetRows(rows)
}

func (m *MainModel) removeRow(id int64) {
	for i, proj := range m.projects {
		if proj.ID == id {
			m.projects = append(m.projects[:i], m.projects[i+1:]...)
			return
		}
	}
}

func InitialModel(app *App) MainModel {
	initalProjects, err := app.Q.GetProjects(app.Ctx)
	if err != nil {
		log.Fatal(err)
	}

	columns := []table.Column{
		{Title: "Id", Width: 4},
		{Title: "Title", Width: 20},
		{Title: "Link", Width: 30},
		{Title: "Icon path", Width: 30},
	}

	m := MainModel{
		app:           app,
		projects:      initalProjects,
		createProject: initialCreateProjectModel(),
		projectTable: table.New(
			table.WithColumns(columns),
			table.WithFocused(true),
			table.WithHeight(7),
			table.WithWidth(85),
		),
	}

	m.updateRows()

	s := table.DefaultStyles()
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	m.projectTable.SetStyles(s)
	return m
}
