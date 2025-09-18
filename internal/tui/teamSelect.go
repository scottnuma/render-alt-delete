package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) viewTeamSelect(add func(...string)) {
	add("What workspace should we delete services from?\n\n")

	for i, owner := range m.owners {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		add(fmt.Sprintf("%s %s - %s\n", cursor, owner.Name, owner.Email))
	}
}

func (m *model) initTeamSelect() {
	m.status = statusSelectTeam
	m.cursor = 0
	owners, err := m.renderSvc.ListAuthorizedOwners()
	if err != nil {
		m.errMsg = fmt.Sprintf("failed to list authorized owners: %s", err)
		m.status = statusError
		return
	}
	m.owners = owners
}

func (m *model) updateKeyMsgSelectTeam(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}

	case "down", "j":
		if m.cursor < len(m.owners)-1 {
			m.cursor++
		}

	case "enter", " ":
		m.ownerID = m.owners[m.cursor].ID
		m.initSelect()
	}
	return m, nil
}
