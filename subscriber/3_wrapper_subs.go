package subscriber

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
)

func OnWrapperCompleted(event *events.WrapperEvent) {
	if event.Status != events.Completed {
		return
	}

	fmt.Printf("[OnWrapperCompleted][%s][%s]", event.Method, event.Status)
	fmt.Println("")
}
