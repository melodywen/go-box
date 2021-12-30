package events

import (
	"fmt"
	"github.com/melodywen/go-box/collections"
	contractsEvents "github.com/melodywen/go-box/contracts/events"
	"github.com/melodywen/go-box/support"
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

// HasListeners Determine if a given event has listeners.
func (dispatcher *Dispatcher) HasListeners(event interface{}) bool {
	eventIndex := dispatcher.app.AbstractToString(event)
	if _, ok := dispatcher.listeners[eventIndex]; ok {
		return true
	}
	if _, ok := dispatcher.wildcards[eventIndex]; ok {
		return true
	}
	return dispatcher.hasWildcardListeners(eventIndex)
}

// hasWildcardListeners Determine if the given event has any wildcard listeners.
func (dispatcher *Dispatcher) hasWildcardListeners(event string) bool {
	for index, _ := range dispatcher.wildcards {
		if (support.StrFun{}).Is(index, event) {
			return true
		}
	}
	return false
}

// GetListeners Get all of the listeners for a given event name.
func (dispatcher *Dispatcher) getWildcardListeners() {
	panic("implement me")
}

// GetListeners Get all of the listeners for a given event name.
func (dispatcher *Dispatcher) GetListeners() {
	panic("implement me")
}

func (dispatcher *Dispatcher) Subscribe() {
	panic("implement me")
}

func (dispatcher *Dispatcher) Until() {
	panic("implement me")
}

// Dispatch Fire an event and call the listeners.
func (dispatcher *Dispatcher) Dispatch(event interface{}, payload interface{}, halt bool) []interface{} {
	// When the given "event" is actually an object we will assume it is an event
	// object and use the class as the event name and this event itself as the
	// payload to the handler, which makes object based events quite simple.
	var wrapPayload []interface{}
	event, wrapPayload = dispatcher.parseEventAndPayload(event, payload)

	// TODO  broadcastable event

	responses := make([]interface{}, 0)

	fmt.Println(responses)

	fmt.Println(212313221, event, wrapPayload)

	return []interface{}{1, 2, 3}
}

// Parse the given event and payload and prepare them for dispatching.
func (dispatcher *Dispatcher) parseEventAndPayload(event interface{}, payload interface{}) (eventOut interface{}, wrapPayload []interface{}) {
	kind := reflect.TypeOf(event).Kind()
	eventOut = dispatcher.app.AbstractToString(event)
	if kind == reflect.Ptr || kind == reflect.Struct {
		wrapPayload = []interface{}{event}
	} else {
		wrapPayload = collections.ArrFun{}.Wrap(payload)
	}
	return eventOut, wrapPayload
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
