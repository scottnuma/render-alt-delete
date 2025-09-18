package tui

import (
	"fmt"
)

func (m model) viewDeleting(add func(...string)) {
	add("\n")
	for _, res := range m.resources {
		if !res.selected {
			continue
		}
		add(" ", res.deleteStatus, " ", res.name, " - ", res.resourceType, "\n")
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
