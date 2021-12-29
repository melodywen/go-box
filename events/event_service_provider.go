package events

import (
	"fmt"
	"github.com/melodywen/go-box/contracts/foundation"
	"github.com/melodywen/go-box/support"
)

type EventServiceProvider struct {
	support.ServiceProvider
}

func NewEventServiceProvider(app foundation.ApplicationInterface) *EventServiceProvider {
	return &EventServiceProvider{ServiceProvider: *support.NewServiceProvider(app)}
}

func (provider *EventServiceProvider) Boot() {

}

func (provider *EventServiceProvider) Register() {
	provider.App.Singleton("events", func(app foundation.ApplicationInterface) {
		fmt.Println("这是一个 event 的具体实现")
	})
}
