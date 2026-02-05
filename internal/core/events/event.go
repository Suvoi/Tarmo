package events

type DomainEvent interface {
	EventName() string
}

type EventBus interface {
	Publish(event DomainEvent)
}
