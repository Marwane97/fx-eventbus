package subscriber

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
)

func OnManagerFailed(event *events.ManagerEvent) {
	if event.Status != events.Failed {
		return
	}

	fmt.Printf("[OnManagerFailed][%s][%s]", event.Method, event.Status)
	fmt.Println("")
}
