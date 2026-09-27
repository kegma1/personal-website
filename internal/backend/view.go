package backend

import (
	tea "charm.land/bubbletea/v2"
)

func (m MainModel) View() tea.View {
	s := ""
	if len(m.projects) > 0 {
		s += m.projectTable.View() + "\n  " + m.projectTable.HelpView() + "\n"
	} else {
		s += "There are no projects in the database..."
	}

	if m.err != nil {
		s += m.err.Error()
	}

	v := tea.NewView(s)

	if m.createProject.open {
		v = m.createProject.View()
	}

	v.AltScreen = true
	return v
}
