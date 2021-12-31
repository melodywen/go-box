package foundation

import (
	events2 "github.com/melodywen/go-box/contracts/events"
	"github.com/melodywen/go-box/contracts/foundation"
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
