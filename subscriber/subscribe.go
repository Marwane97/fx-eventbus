package subscriber

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
	"github.com/asaskevich/EventBus"
)

func New() EventBus.Bus {
	return EventBus.New()
}

func RegisterSubscribers(bus EventBus.Bus) error {
	if err := bus.Subscribe(string(events.ServiceName), OnServiceStarted); err != nil {
		return fmt.Errorf("error subscribing to OnServiceStarted: %v", err)
	}

	if err := bus.Subscribe(string(events.ServiceName), OnServiceFailed); err != nil {
		return fmt.Errorf("error subscribing to OnServiceFailed: %v", err)
	}

	if err := bus.Subscribe(string(events.ServiceName), OnServiceCompleted); err != nil {
		return fmt.Errorf("error subscribing to OnServiceCompleted: %v", err)
	}

	if err := bus.Subscribe(string(events.ManagerName), OnManagerStarted); err != nil {
		return fmt.Errorf("error subscribing to OnManagerStarted: %v", err)
	}

	if err := bus.Subscribe(string(events.ManagerName), OnManagerFailed); err != nil {
		return fmt.Errorf("error subscribing to OnManagerFailed: %v", err)
	}

	if err := bus.Subscribe(string(events.ManagerName), OnManagerCompleted); err != nil {
		return fmt.Errorf("error subscribing to OnManagerCompleted: %v", err)
	}

	if err := bus.Subscribe(string(events.WrapperName), OnWrapperStarted); err != nil {
		return fmt.Errorf("error subscribing to OnWrapperStarted: %v", err)
	}

	if err := bus.Subscribe(string(events.WrapperName), OnWrapperFailed); err != nil {
		return fmt.Errorf("error subscribing to OnWrapperFailed: %v", err)
	}

	if err := bus.Subscribe(string(events.WrapperName), OnWrapperCompleted); err != nil {
		return fmt.Errorf("error subscribing to OnWrapperCompleted: %v", err)
	}

	return nil
}
