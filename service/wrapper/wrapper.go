package wrapper

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
	"github.com/asaskevich/EventBus"
)

type Wrapper struct {
	bus EventBus.Bus
}

func New(bus EventBus.Bus) *Wrapper {
	return &Wrapper{
		bus: bus,
	}
}

func (w Wrapper) DoSomeThing(isError bool) error {
	wrapperEvent := events.NewWrapperEvent(events.WrapperName, events.Started, "search")

	w.bus.Publish(string(events.WrapperName), wrapperEvent)

	if isError {
		wrapperEvent = events.NewWrapperEvent(events.WrapperName, events.Failed, "search")
		w.bus.Publish(string(events.WrapperName), wrapperEvent)

		return fmt.Errorf("an error occur ")
	}

	wrapperEvent = events.NewWrapperEvent(events.WrapperName, events.Completed, "search")
	w.bus.Publish(string(events.WrapperName), wrapperEvent)

	return nil
}
