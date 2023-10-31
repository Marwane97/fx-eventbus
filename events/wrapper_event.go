package events

type WrapperEvent struct {
	Name   EventName
	Status EventStatus
	Method string
	Flight map[string]interface{}
}

func NewWrapperEvent(name EventName, status EventStatus, method string) *WrapperEvent {
	return &WrapperEvent{
		Name:   name,
		Status: status,
		Method: method,
		Flight: nil,
	}
}
