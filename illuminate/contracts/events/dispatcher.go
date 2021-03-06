package events

// ListenerFun listener function contract
type ListenerFun func(args ...interface{}) interface{}

// WrapListenerFun warp listener function contract
type WrapListenerFun func(event string, payload []interface{}) interface{}

// DispatcherInterface  dispatcher contract
type DispatcherInterface interface {
	// Listen Register an event listener with the dispatcher.
	Listen(events interface{}, listener ListenerFun)
	// HasListeners Determine if a given event has listeners.
	HasListeners(event interface{}) bool
	// Subscribe Register an event subscriber with the dispatcher.
	Subscribe()
	// Until Dispatch an event until the first non-null response is returned.
	Until(event interface{}, payload interface{}) interface{}
	// Dispatch an event and call the listeners.
	Dispatch(event interface{}, payload interface{}, halt bool) interface{}
	// Push Register an event and payload to be fired later.
	Push(event interface{}, payload interface{})
	// Flush a set of pushed events.
	Flush()
	// Forget Remove a set of listeners from the dispatcher.
	Forget()
	// ForgetPushed Forget all of the queued listeners.
	ForgetPushed()
}
