package subscriber

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
)

func OnManagerCompleted(event *events.ManagerEvent) {
	if event.Status != events.Completed {
		return
	}

	fmt.Printf("[OnManagerCompleted][%s][%s]", event.Method, event.Status)
	fmt.Println("")
}
