package subscriber

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
)

func OnServiceFailed(event *events.ServiceEvent) {
	if event.Status != events.Failed {
		return
	}

	fmt.Printf("[OnServiceFailed][%s][%s]", event.Method, event.Status)
	fmt.Println("")
}
