package tui

import (
	"fmt"

	"github.com/scottnuma/render-alt-delete/internal/log"
)

func (m model) viewDeleting(add func(...string)) {
	add("\n")
	for _, res := range m.resources {
		if !res.selected {
			continue
		}
		add(" ", res.deleteStatus, " ", res.name, "\n")
	}
}

func (m *model) initDeleting() {
	m.status = statusDeleting
	m.cursor = 0
	go func() {
		for _, resInfo := range m.resources {
			if resInfo.selected && resInfo.deleteStatus == deleteStatusPending {
				resInfo.deleteStatus = deleteStatusWorking
				m.deleteStatusUpdate <- struct{}{}

				err := resInfo.delete(m.renderSvc)
				if err != nil {
					log.Logger.Error("failed to delete", "err", err)
					resInfo.deleteStatus = fmt.Sprintf(
						"failed to delete: %s",
						err,
					)
				} else {
					resInfo.deleteStatus = deleteStatusDone
				}

				m.deleteStatusUpdate <- struct{}{}
			}
		}
	}()
}
