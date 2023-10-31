package subscriber

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
)

func OnManagerStarted(event *events.ManagerEvent) {
	if event.Status != events.Started {
		return
	}

	fmt.Printf("[OnManagerStarted][%s][%s]", event.Method, event.Status)
	fmt.Println("")
}
