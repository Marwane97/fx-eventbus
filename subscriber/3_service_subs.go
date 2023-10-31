package subscriber

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
)

func OnServiceCompleted(event *events.ServiceEvent) {
	if event.Status != events.Completed {
		return
	}

	fmt.Printf("[OnServiceCompleted][%s][%s]", event.Method, event.Status)
	fmt.Println("###############  ###############  ###############")
	fmt.Println("")
}
