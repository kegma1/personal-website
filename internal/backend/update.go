package backend

import (
	"fmt"
	"net/url"
	"personal_website/repositories"
	"strconv"

	tea "charm.land/bubbletea/v2"
)

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	if m.createProject.open {
		var cmd tea.Cmd

		m.createProject, cmd = m.createProject.Update(msg)
		cmds = append(cmds, cmd)

		return m, tea.Batch(cmds...)

	}

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "c":
			m.createProject = initialCreateProjectModel()
			m.createProject.open = true
			return m, nil

		case "d":
			selectedID, err := strconv.Atoi(m.projectTable.SelectedRow()[0])
			if err != nil { // should in theroy never happen
				m.err = err
				return m, nil
			}

			m.app.Q.DeleteProject(m.app.Ctx, int64(selectedID))
			m.removeRow(int64(selectedID))
			m.updateRows()
		}

	case SubmitProjectMsg:
		title := msg.Title
		link := msg.Link
		_, err := url.ParseRequestURI(link)
		if err != nil {
			m.err = fmt.Errorf("invalid URL: %w", err)
			return m, nil
		}

		proj, err := m.app.Q.CreateProject(m.app.Ctx, repositories.CreateProjectParams{
			Title: title,
			Link:  link,
		})
		if err != nil {
			m.err = err
			return m, nil
		}

		m.err = nil
		m.projects = append(m.projects, proj)
		m.updateRows()
	}

	var cmd tea.Cmd

	m.projectTable, cmd = m.projectTable.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
