package foundation

import (
	events2 "github.com/melodywen/go-box/illuminate/contracts/events"
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	log2 "github.com/melodywen/go-box/illuminate/contracts/log"
)

// HasBeenBootstrapped Determine if the application has been bootstrapped before.
func (app *Application) HasBeenBootstrapped() bool {
	return app.hasBeenBootstrapped
}

// BootstrapWith Run the given array of bootstrap classes.
func (app *Application) BootstrapWith(bootstrappers []foundation.BootstrapInterface) {
	app.hasBeenBootstrapped = true
	var dispatcher events2.DispatcherInterface
	dispatcher = app.Make("events").(events2.DispatcherInterface)
	for _, bootstrapper := range bootstrappers {
		dispatcher.Dispatch("bootstrapping:"+app.AbstractToString(bootstrapper), app, false)
		bootstrapper.Bootstrap(app)
		dispatcher.Dispatch("bootstrapped:"+app.AbstractToString(bootstrapper), app, false)
	}
}

// BootstrapOpenListen bootstrap open listened
func (app *Application) BootstrapOpenListen() {
	var dispatcher events2.DispatcherInterface
	dispatcher = app.Make("events").(events2.DispatcherInterface)
	var log log2.LoggerInterface
	log = app.Make("log").(log2.LoggerInterface)
	dispatcher.Listen("bootstrapping:*", func(args ...interface{}) interface{} {
		log.Info(args[0].(string), nil)
		return nil
	})
	dispatcher.Listen("bootstrapped:*", func(args ...interface{}) interface{} {
		log.Info(args[0].(string), nil)
		return nil
	})
}
