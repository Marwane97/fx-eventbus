package subscriber

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
)

func OnWrapperFailed(event *events.WrapperEvent) {
	if event.Status != events.Failed {
		return
	}

	fmt.Printf("[OnWrapperFailed][%s][%s]", event.Method, event.Status)
	fmt.Println("")
}
