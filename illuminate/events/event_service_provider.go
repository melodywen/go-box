package events

import (
	"github.com/melodywen/go-box/illuminate/contracts/events"
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	"github.com/melodywen/go-box/illuminate/support"
)

// EventServiceProvider struct
type EventServiceProvider struct {
	support.ServiceProvider
}

// NewEventServiceProvider new instance
func NewEventServiceProvider(app foundation.ApplicationInterface) *EventServiceProvider {
	return &EventServiceProvider{ServiceProvider: *support.NewServiceProvider(app)}
}

// Register rewrite method
func (provider *EventServiceProvider) Register() {
	provider.App.Singleton("events", func(app foundation.ApplicationInterface) events.DispatcherInterface {
		return NewDispatcher(app)
	})
}
