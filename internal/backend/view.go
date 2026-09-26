package backend

import (
	"fmt"

	tea "charm.land/bubbletea/v2"
)

func (m Model) View() tea.View {
	s := ""
	if len(m.projects) > 0 {
		for _, project := range m.projects {
			s += fmt.Sprintf("%s %s\n", project.Title, project.Link)
		}
	} else {
		s += "There are no projects in the database..."
	}
	return tea.NewView(s)
}
