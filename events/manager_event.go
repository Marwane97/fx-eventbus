package events

type ManagerEvent struct {
	Name   EventName
	Status EventStatus
	Method string
}

func NewManagerEvent(name EventName, status EventStatus, method string) *ManagerEvent {
	return &ManagerEvent{
		Name:   name,
		Status: status,
		Method: method,
	}
}
