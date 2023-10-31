package subscriber

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
)

func OnWrapperStarted(event *events.WrapperEvent) {
	if event.Status != events.Started {
		return
	}

	fmt.Printf("[OnWrapperStarted][%s][%s]", event.Method, event.Status)
	fmt.Println("")
}
