package domain

type EventsDB interface {
	CreateEvent(event Event)
}
