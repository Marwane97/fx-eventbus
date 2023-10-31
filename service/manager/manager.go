package manager

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
	"github.com/Marwane97/fx-eventbus/service/wrapper"
	"github.com/asaskevich/EventBus"
)

type Manager struct {
	bus     EventBus.Bus
	wrapper *wrapper.Wrapper
}

func New(bus EventBus.Bus) *Manager {
	return &Manager{
		bus:     bus,
		wrapper: wrapper.New(bus),
	}
}

func (m *Manager) DoSomeThing(isError bool) error {
	managerEvent := events.NewManagerEvent(events.ManagerName, events.Started, "search")

	m.bus.Publish(string(events.ManagerName), managerEvent)

	err := m.wrapper.DoSomeThing(isError)

	if err != nil {
		managerEvent = events.NewManagerEvent(events.ManagerName, events.Failed, "search")
		m.bus.Publish(string(events.ManagerName), managerEvent)
		return fmt.Errorf("an error occur ")
	}

	managerEvent = events.NewManagerEvent(events.ManagerName, events.Completed, "search")
	m.bus.Publish(string(events.ManagerName), managerEvent)

	return nil
}
