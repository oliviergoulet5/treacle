package tui

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type model struct {
	ready  bool
	width  int
	height int
}

func New() *tea.Program {
	return tea.NewProgram(model{})
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.ready = true
		m.width = msg.Width
		m.height = msg.Height
		_ = msg
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() tea.View {
	v := tea.NewView("Loading...")
	if m.ready {
		v.SetContent(lipgloss.NewStyle().
			Width(m.width).
			Height(m.height).
			Padding(1, 2).
			Align(lipgloss.Top, lipgloss.Left).
			Background(lipgloss.Color("#000000")).
			Render("treacle"))
	}
	v.AltScreen = true
	return v
}
