package events

type EventName string

const (
	ServiceName EventName = "service.event"
	WrapperName EventName = "wrapper.event"
	ManagerName EventName = "manager.event"
)
