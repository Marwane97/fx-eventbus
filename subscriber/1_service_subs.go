package subscriber

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
)

func OnServiceStarted(event *events.ServiceEvent) {
	if event.Status != events.Started {
		return
	}

	fmt.Printf("[OnServiceStarted][%s][%s]", event.Method, event.Status)
	fmt.Println("")
}
