package foundation

import (
	"fmt"
	"github.com/melodywen/go-box/illuminate/contracts/support"
)

// Boot the application's service providers.
func (app *Application) Boot() {
	if app.IsBooted() {
		return
	}

	// Once the application has booted we will also fire some "booted" callbacks
	// for any listeners that need to do work after this initial booting gets
	// finished. This is useful when ordering the boot-up processes we run.
}

// RegisterConfiguredProviders Register all the configured providers.
func (app *Application) RegisterConfiguredProviders() {
	NewProviderRepository(app).Load()
}

// AddDeferredServices Add an array of services to the application's deferred services.
func (app *Application) AddDeferredServices(services map[string]support.ServiceProviderInterface) {
	for serviceName, provider := range services {
		app.deferredServices[serviceName] = provider
	}
}

// ResolveCallback resolve the given type from the container.
func (app *Application) ResolveCallback(abstract string)  {
	fmt.Println(1231312,abstract)

}
