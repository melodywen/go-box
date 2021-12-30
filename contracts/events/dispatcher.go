package events

type ListenerFun func(args ...interface{}) interface{}
type WrapListenerFun func(event string, payload []interface{}) interface{}

type DispatcherInterface interface {
	// Listen Register an event listener with the dispatcher.
	Listen(events interface{}, listener ListenerFun)
	// HasListeners Determine if a given event has listeners.
	HasListeners(event interface{}) bool
	// Subscribe Register an event subscriber with the dispatcher.
	Subscribe()
	// Until Dispatch an event until the first non-null response is returned.
	Until()
	// Dispatch an event and call the listeners.
	Dispatch(event interface{}, payload interface{}, halt bool) interface{}
	// Push Register an event and payload to be fired later.
	Push()
	// Flush a set of pushed events.
	Flush()
	// Forget Remove a set of listeners from the dispatcher.
	Forget()
	// ForgetPushed Forget all of the queued listeners.
	ForgetPushed()
}
