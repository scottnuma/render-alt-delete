package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

const (
	deleteStatusPending = "pending"
	deleteStatusWorking = "working"
	deleteStatusDone    = "deleted"
)

func (m model) viewReview(add func(...string)) {
	add("Going to delete the following services:\n\n")
	for _, res := range m.resources {
		if !res.selected {
			continue
		}
		add(" ", res.name, "\n")
	}

	if m.cursor == 0 {
		add(
			"\n",
			" [x] No\n",
			" [ ] Yes\n",
		)
	} else {
		add(
			"\n",
			" [ ] No\n",
			" [x] Yes\n",
		)
	}
}

func (m *model) initReview() {
	m.status = statusReview
	m.cursor = 0

	for _, res := range m.resources {
		if !res.selected {
			continue
		}
		res.deleteStatus = deleteStatusPending
	}
}

func (m *model) updateKeyMsgReview(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}

	case "down", "j":
		if m.cursor < 1 {
			m.cursor++
		}
	case "enter", " ":
		if m.cursor == 1 {
			m.initDeleting()
		} else {
			m.status = statusSelect
		}
	}

	return m, nil
}
