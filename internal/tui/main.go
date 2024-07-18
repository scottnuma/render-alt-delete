package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/scottnuma/render-alt-delete/internal/rad"
)

type RenderService interface {
	ListServices(ownerID string) ([]rad.Service, error)
	DeleteService(serviceID string) error
	ListPostgres(ownerID string) ([]rad.Postgres, error)
	DeletePostgres(postgresID string) error
	ListRedis(ownerID string) ([]rad.Redis, error)
	DeleteRedis(redisID string) error
	ListAuthorizedOwners() ([]rad.Owner, error)
}

func NewTUI(renderSvc RenderService) *tea.Program {
	return tea.NewProgram(newModel(renderSvc))

}

type status int

const (
	statusSelectTeam = iota + 1
	statusSelect
	statusReview
	statusDeleting
)

type model struct {
	status             status
	resources          []*resource
	cursor             int
	deleteStatusUpdate chan struct{}
	renderSvc          RenderService
	ownerID            string
	owners             []rad.Owner
}

type resource struct {
	name         string
	resourceType string
	deleteStatus string
	selected     bool
	delete       func(RenderService) error
}

func newModel(renderSvc RenderService) model {
	m := model{
		deleteStatusUpdate: make(chan struct{}),
		renderSvc:          renderSvc,
	}
	m.initTeamSelect()
	return m
}

func (m model) Init() tea.Cmd {
	return waitForUpdate(m.deleteStatusUpdate)

}

type responseMsg struct{}

func waitForUpdate(sub chan struct{}) tea.Cmd {
	return func() tea.Msg {
		return responseMsg(<-sub)
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case responseMsg:
		return m, waitForUpdate(m.deleteStatusUpdate)

	case tea.KeyMsg:
		// Global Key Presses
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}

		switch m.status {
		case statusReview:
			return m.updateKeyMsgReview(msg)
		case statusSelect:
			return m.updateKeyMsgSelect(msg)
		case statusSelectTeam:
			return m.updateKeyMsgSelectTeam(msg)
		}
	}

	return m, nil
}

// concatenate returns the concatendated strings provided to the add func.
func concatenate(
	creator func(
		add func(...string),
	),
) string {
	screen := []string{}
	add := func(line ...string) { screen = append(screen, line...) }
	creator(add)
	return strings.Join(screen, "")
}

func (m model) View() string {
	return m.styleContent(concatenate(m.getViewContent))
}

func (m model) styleContent(content string) string {
	return baseStyle.Render(content)
}

func (m model) getViewContent(add func(...string)) {
	switch m.status {
	case statusSelectTeam:
		add(concatenate(m.viewTeamSelect))
	case statusSelect:
		add(concatenate(m.viewSelect))
	case statusReview:
		add(concatenate(m.viewReview))
	case statusDeleting:
		add(concatenate(m.viewDeleting))
	default:
		add(fmt.Sprintf("status %v not implemented", m.status))
	}

	// Global footer
	add("\nPress q to quit.")
}
