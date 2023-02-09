package tui

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *model) updateKeyMsgSelect(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}

	case "down", "j":
		if m.cursor < len(m.serviceInfos) {
			m.cursor++
		}

	case "enter", " ":
		if m.cursor == len(m.serviceInfos) {
			m.initReview()
		} else {
			m.serviceInfos[m.cursor].selected = !m.serviceInfos[m.cursor].selected
		}
	}
	return m, nil
}
func (m model) viewSelect(add func(...string)) {
	add("What services should we delete?\n\n")

	for i, svcInfo := range m.serviceInfos {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if svcInfo.selected {
			checked = "x"
		}

		add(fmt.Sprintf("%s [%s] %s\n", cursor, checked, svcInfo.svc.Name))
	}

	cursor := " "
	if m.cursor == len(m.serviceInfos) {
		cursor = ">"
	}
	add(fmt.Sprintf("\n%s Done\n", cursor))
}

func (m *model) initSelect() {
	m.status = statusSelect
	m.cursor = 0

	svcs, err := m.renderSvc.ListServices(m.ownerID)
	if err != nil {
		log.Println("failed to list services: ", err)
		os.Exit(1)
	}

	m.serviceInfos = make([]*serviceInfo, len(svcs))
	for i, svc := range svcs {
		m.serviceInfos[i] = &serviceInfo{svc: svc}
	}
}
