package events

type EventStatus string

const (
	Failed    EventStatus = "failed"
	Started   EventStatus = "started"
	Completed EventStatus = "completed"
)
