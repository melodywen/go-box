package events

import (
	contractsEvents "github.com/melodywen/go-box/contracts/events"
	"github.com/melodywen/go-ioc/contracts"
	"reflect"
	"strings"
)

type Dispatcher struct {
	app            contracts.ContainerContract
	listeners      map[string][]contractsEvents.WrapListenerFun
	wildcards      map[string][]contractsEvents.WrapListenerFun
	wildcardsCache []string
	queueResolver  interface{}
}

func NewDispatcher(app contracts.ContainerContract) *Dispatcher {
	return &Dispatcher{
		app:            app,
		listeners:      map[string][]contractsEvents.WrapListenerFun{},
		wildcards:      map[string][]contractsEvents.WrapListenerFun{},
		wildcardsCache: nil,
		queueResolver:  nil,
	}
}

// Listen Register an event listener with the dispatcher.
func (dispatcher *Dispatcher) Listen(events interface{}, listener contractsEvents.ListenerFun) {
	eventsIndex := make([]string, 0)
	switch reflect.TypeOf(events).Kind() {
	case reflect.Slice:
		if value, ok := events.([]interface{}); ok {
			for _, v := range value {
				eventsIndex = append(eventsIndex, dispatcher.app.AbstractToString(v))
			}
		}
	default:
		eventsIndex = append(eventsIndex, dispatcher.app.AbstractToString(events))
	}

	for _, index := range eventsIndex {
		if strings.Contains(index, "*") {
			dispatcher.setupWildcardListen(index, listener)
		} else {
			if dispatcher.listeners[index] == nil {
				dispatcher.listeners[index] = make([]contractsEvents.WrapListenerFun, 0)
			}
			dispatcher.listeners[index] = append(dispatcher.listeners[index], dispatcher.MakeListener(listener, false))
		}
	}
}

// set up a wildcard listener callback.
func (dispatcher *Dispatcher) setupWildcardListen(event string, listener contractsEvents.ListenerFun) {
	if dispatcher.wildcards[event] == nil {
		dispatcher.wildcards[event] = make([]contractsEvents.WrapListenerFun, 0)
	}
	dispatcher.wildcards[event] = append(dispatcher.wildcards[event], dispatcher.MakeListener(listener, true))
	dispatcher.wildcardsCache = []string{}
}

// MakeListener Register an event listener with the dispatcher.
func (dispatcher *Dispatcher) MakeListener(listener contractsEvents.ListenerFun, wildcard bool) contractsEvents.WrapListenerFun {
	return func(event string, payload []interface{}) interface{} {
		if wildcard {
			return listener(event, payload)
		}
		return listener(payload...)
	}
}

func (dispatcher *Dispatcher) HasListeners() {
	panic("implement me")
}

func (dispatcher *Dispatcher) Subscribe() {
	panic("implement me")
}

func (dispatcher *Dispatcher) Until() {
	panic("implement me")
}

func (dispatcher *Dispatcher) Dispatch() {
	panic("implement me")
}

func (dispatcher *Dispatcher) Push() {
	panic("implement me")
}

func (dispatcher *Dispatcher) Flush() {
	panic("implement me")
}

func (dispatcher *Dispatcher) Forget() {
	panic("implement me")
}

func (dispatcher *Dispatcher) ForgetPushed() {
	panic("implement me")
}
