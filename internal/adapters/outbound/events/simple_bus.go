package events

import "tarmo/internal/core/events"

type SimpleEventBus struct {
	listeners []func(events.DomainEvent)
}

func (b *SimpleEventBus) Publish(e events.DomainEvent) {
	for _, listener := range b.listeners {
		listener(e) // Síncrono (o lanza goroutine si quieres async)
	}
}
func (b *SimpleEventBus) Subscribe(fn func(events.DomainEvent)) {
	b.listeners = append(b.listeners, fn)
}
