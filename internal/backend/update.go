package backend

import (
	"log"
	"net/url"
	"personal_website/repositories"

	tea "charm.land/bubbletea/v2"
)

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	if !m.createProject.open {
		switch msg := msg.(type) {
		case tea.KeyPressMsg:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "c":
				m.createProject = initialCreateProjectModel()
				m.createProject.open = true
				return m, nil
			}
		case SubmitProjectMsg:
			title := msg.Title
			link := msg.Link
			_, err := url.ParseRequestURI(link)
			if err != nil {
				log.Fatal(err)
			}

			proj, err := m.app.Q.CreateProject(m.app.Ctx, repositories.CreateProjectParams{
				Title: title,
				Link: link,
			})
			if err != nil {
				log.Fatal(err)
			}

			m.projects = append(m.projects, proj)

		}
	}

	m.createProject, cmd = m.createProject.Update(msg)
	return m, cmd
}
