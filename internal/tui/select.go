package tui

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/scottnuma/render-alt-delete/internal/log"
)

func (m *model) updateKeyMsgSelect(msg tea.KeyMsg) (tea.Model, tea.Cmd) {

	switch msg.String() {
	case "up", "k":
		if m.cursor > 0 {
			m.cursor--
		}

	case "down", "j":
		if m.cursor < len(m.resources) {
			m.cursor++
		}

	case "enter", " ":
		if m.cursor == len(m.resources) {
			m.initReview()
		} else {
			m.resources[m.cursor].selected = !m.resources[m.cursor].selected
		}
	}
	return m, nil
}
func (m model) viewSelect(add func(...string)) {
	add("What services should we delete?\n\n")

	for i, resInfo := range m.resources {

		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if resInfo.selected {
			checked = "x"
		}

		add(fmt.Sprintf("%s [%s] %s\n", cursor, checked, resInfo.name))
	}

	cursor := " "
	if m.cursor == len(m.resources) {
		cursor = ">"
	}
	add(fmt.Sprintf("\n%s Done\n", cursor))
}

func (m *model) initSelect() {
	m.status = statusSelect
	m.cursor = 0

	m.resources = []*resource{}

	svcs, err := m.renderSvc.ListServices(m.ownerID)
	if err != nil {
		log.Logger.Error("failed to list services", "err", err)
		os.Exit(1)
	}

	for _, svc := range svcs {
		m.resources = append(m.resources, &resource{
			name:         svc.Name,
			resourceType: "Service",
			delete: func(renderSvc RenderService) error {
				return renderSvc.DeleteService(svc.ID)
			},
		})
	}

	dbs, err := m.renderSvc.ListPostgres(m.ownerID)
	if err != nil {
		log.Logger.Error("failed to list postgres databases", "err", err)
		os.Exit(1)
	}

	for _, db := range dbs {
		m.resources = append(m.resources, &resource{
			name:         db.Name,
			resourceType: "Postgres",
			delete: func(renderSvc RenderService) error {
				return renderSvc.DeletePostgres(db.ID)
			},
		})
	}

	redisdbs, err := m.renderSvc.ListRedis(m.ownerID)
	if err != nil {
		log.Logger.Error("failed to list redis databases", "err", err)
		os.Exit(1)
	}

	for _, redis := range redisdbs {
		m.resources = append(m.resources, &resource{
			name:         redis.Name,
			resourceType: "Redis",
			delete: func(renderSvc RenderService) error {
				return renderSvc.DeleteRedis(redis.ID)
			},
		})
	}
}
