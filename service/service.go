package service

import (
	"fmt"
	"github.com/Marwane97/fx-eventbus/events"
	"github.com/Marwane97/fx-eventbus/service/manager"
	"github.com/asaskevich/EventBus"
)

type Search struct {
	bus     EventBus.Bus
	manager *manager.Manager
}

func New(bus EventBus.Bus) *Search {
	return &Search{
		bus:     bus,
		manager: manager.New(bus),
	}
}

func (s *Search) Search(isException bool) error {
	s.bus.Publish(
		string(events.ServiceName),
		events.NewServiceEvent(events.Started),
	)

	err := s.manager.DoSomeThing(isException)

	if err != nil {
		s.bus.Publish(
			string(events.ServiceName),
			events.NewServiceEvent(events.Failed),
		)

		return fmt.Errorf("an error occur ")
	}

	s.bus.Publish(
		string(events.ServiceName),
		events.NewServiceEvent(events.Completed),
	)

	return nil
}
