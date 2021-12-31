package events

import (
	"github.com/melodywen/go-box/collections"
	contractsEvents "github.com/melodywen/go-box/contracts/events"
	"github.com/melodywen/go-box/support"
	"github.com/melodywen/go-ioc/contracts"
	"reflect"
	"strings"
)

// Dispatcher dispatcher struct
type Dispatcher struct {
	app            contracts.ContainerContract
	listeners      map[string][]contractsEvents.WrapListenerFun
	wildcards      map[string][]contractsEvents.WrapListenerFun
	wildcardsCache map[string][]contractsEvents.WrapListenerFun
	queueResolver  interface{}
}

// NewDispatcher new dispatcher instance
func NewDispatcher(app contracts.ContainerContract) *Dispatcher {
	return &Dispatcher{
		app:            app,
		listeners:      map[string][]contractsEvents.WrapListenerFun{},
		wildcards:      map[string][]contractsEvents.WrapListenerFun{},
		wildcardsCache: map[string][]contractsEvents.WrapListenerFun{},
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
		// first char cannot "*"
		if strings.Contains(strings.TrimLeft(index, "*"), "*") {
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
	dispatcher.wildcardsCache = map[string][]contractsEvents.WrapListenerFun{}
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
	for index := range dispatcher.wildcards {
		if (support.StrFun{}).Is(index, event) {
			return true
		}
	}
	return false
}

// GetListeners Get all of the listeners for a given event name.
func (dispatcher *Dispatcher) getWildcardListeners(event string) (wildcards []contractsEvents.WrapListenerFun) {
	for index, fun := range dispatcher.wildcards {
		if (support.StrFun{}).Is(index, event) {
			wildcards = append(wildcards, fun...)
		}
	}
	dispatcher.wildcardsCache[event] = wildcards
	return wildcards
}

// GetListeners Get all of the listeners for a given event name.
func (dispatcher *Dispatcher) GetListeners(event interface{}) (response []contractsEvents.WrapListenerFun) {
	eventIndex := dispatcher.app.AbstractToString(event)
	if value, ok := dispatcher.listeners[eventIndex]; ok {
		response = append(response, value...)
	}
	if value, ok := dispatcher.wildcardsCache[eventIndex]; ok {
		response = append(response, value...)
	} else {
		value := dispatcher.getWildcardListeners(eventIndex)
		response = append(response, value...)
	}
	return response
}

// Until Dispatch an event until the first non-null response is returned.
func (dispatcher *Dispatcher) Until(event interface{}, payload interface{}) interface{} {
	return dispatcher.Dispatch(event, payload, true)
}

// Dispatch Fire an event and call the listeners.
func (dispatcher *Dispatcher) Dispatch(event interface{}, payload interface{}, halt bool) interface{} {
	// When the given "event" is actually an object we will assume it is an event
	// object and use the class as the event name and this event itself as the
	// payload to the handler, which makes object based events quite simple.
	var wrapPayload []interface{}
	event, wrapPayload = dispatcher.parseEventAndPayload(event, payload)

	// TODO  broadcastable event

	responses := make([]interface{}, 0)

	for _, listener := range dispatcher.GetListeners(event) {
		response := listener(event.(string), wrapPayload)

		// If a response is returned from the listener and event halting is enabled
		// we will just return this response, and not call the rest of the event
		// listeners. Otherwise we will add the response on the response list.
		if halt && response != nil {
			return response
		}

		// If a boolean false is returned from a listener, we will stop propagating
		// the event to any further listeners down in the chain, else we keep on
		// looping through the listeners and firing every one in our sequence.
		if value, ok := response.(bool); ok && !value {
			break
		}

		responses = append(responses, response)
	}
	if halt {
		return nil
	}
	return responses
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

// Push Register an event and payload to be fired later.
func (dispatcher *Dispatcher) Push(event interface{}, payload interface{}) {
	panic("implement me")
}

// Subscribe Register an event subscriber with the dispatcher.
func (dispatcher *Dispatcher) Subscribe() {
	panic("implement me")
}

// Flush a set of pushed events.
func (dispatcher *Dispatcher) Flush() {
	panic("implement me")
}

// Forget Remove a set of listeners from the dispatcher.
func (dispatcher *Dispatcher) Forget() {
	panic("implement me")
}

// ForgetPushed Forget all of the queued listeners.
func (dispatcher *Dispatcher) ForgetPushed() {
	panic("implement me")
}
