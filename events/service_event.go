package events

type ServiceEvent struct {
	Name         EventName
	Method       string
	Status       EventStatus
	SearchParams map[string]interface{}
}

func NewServiceEvent(status EventStatus) *ServiceEvent {
	return &ServiceEvent{
		Name:         ServiceName,
		Method:       "Search",
		Status:       status,
		SearchParams: map[string]interface{}{"request_id": "2345567689"},
	}
}
