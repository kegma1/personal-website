package backend

import (
	"strings"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
)

const (
	titleFocused = iota
	linkFocused
)


type CreateProjectModel struct {
	open    bool
	inputs  []textinput.Model
	focused int
}

type SubmitProjectMsg struct {
	Title string
	Link  string
} 

func submitProject(title, link string) tea.Cmd {
	return func() tea.Msg {
		return SubmitProjectMsg{
			Title: title,
			Link: link,
		}
	}
}

func initialCreateProjectModel() CreateProjectModel {
	// 0_0
	cp := CreateProjectModel{
		open:    false,
		inputs:  make([]textinput.Model, 2),
		focused: titleFocused,
	}

	var t textinput.Model
	for i := range cp.inputs {
		t = textinput.New()

		t.SetWidth(80)
		t.SetVirtualCursor(false)

		switch i {
		case 0:
			t.Placeholder = "Enter title..."
			t.Focus()
		case 1:
			t.Placeholder = "Enter link..."
		}

		cp.inputs[i] = t
	}

	return cp
}

func (m CreateProjectModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m CreateProjectModel) Update(msg tea.Msg) (CreateProjectModel, tea.Cmd) {
	if !m.open {
		return m, nil
	}

	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focused == len(m.inputs)-1 {
				m.open = false
				return m, submitProject(m.inputs[0].Value(), m.inputs[1].Value())
			}

			if s == "up" || s == "shift+tab" {
				m.focused--
			} else {
				m.focused++
			}

			if m.focused > len(m.inputs)-1 {
				m.focused = 0
			} else if m.focused < 0 {
				m.focused = len(m.inputs) - 1
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				if i == m.focused {
					cmds[i] = m.inputs[i].Focus()
					continue
				}
				m.inputs[i].Blur()
			}

			return m, tea.Batch(cmds...)

		case "esc":
			m.open = false
			return m, nil
		}
	}

	cmd = m.updateInputs(msg)

	return m, cmd
}

func (m *CreateProjectModel) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}

func createHeader() string {
	return "Create project"
}

func createFooter() string {
	return "Enter: create • Tab: switch • Esc: cancel"
}

func (m CreateProjectModel) View() tea.View {
	if !m.open {
		return tea.NewView("")
	}

	var c *tea.Cursor
	var b strings.Builder

	b.WriteString(createHeader())
	b.WriteRune('\n')
	for i, in := range m.inputs {
		b.WriteString(in.View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}

		if in.Focused() {
			c = in.Cursor()
			if c != nil {
				c.Y += i
			}
		}
	}

	c.Y++

	b.WriteRune('\n')
	b.WriteString(createFooter())

	v := tea.NewView(b.String())
	v.Cursor = c
	return v
}
