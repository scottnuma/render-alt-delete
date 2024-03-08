package tui

import (
	"fmt"
)

func (m model) viewDeleting(add func(...string)) {
	add("\n")
	for _, svcInfo := range m.serviceInfos {
		if !svcInfo.selected {
			continue
		}
		add(" ", svcInfo.deleteStatus, " ", svcInfo.svc.Name, "\n")
	}
}

func (m *model) initDeleting() {
	m.status = statusDeleting
	m.cursor = 0
	go func() {
		for _, svcInfo := range m.serviceInfos {
			if svcInfo.selected && svcInfo.deleteStatus == deleteStatusPending {
				svcInfo.deleteStatus = deleteStatusWorking
				m.deleteStatusUpdate <- struct{}{}

				err := m.renderSvc.DeleteService(svcInfo.svc.ID)
				if err != nil {
					svcInfo.deleteStatus = fmt.Sprint("failed to delete: %s", err)
				} else {
					svcInfo.deleteStatus = deleteStatusDone
				}

				m.deleteStatusUpdate <- struct{}{}
			}
		}
	}()
}
