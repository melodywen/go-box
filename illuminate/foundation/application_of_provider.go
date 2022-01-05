package foundation

import (
	"fmt"
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
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
	app.fireAppCallbacks(app.bootingCallbacks)

	for _, provider := range app.serviceProviders {
		app.bootProvider(provider)
	}
	app.booted = true

	app.fireAppCallbacks(app.bootedCallbacks)
}

// fireAppCallbacks Call the booting callbacks for the application.
func (app *Application) fireAppCallbacks(callbacks []func(app foundation.ApplicationInterface)) {
	for _, callback := range callbacks {
		callback(app)
	}
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
func (app *Application) ResolveCallback(abstract string) {
	fmt.Println("resolve callback", abstract)
	app.loadDeferredProviderIfNeeded(app.GetAlias(abstract))
}

// Load the deferred provider if the given type is a deferred service and the instance has not been loaded.
func (app *Application) loadDeferredProviderIfNeeded(abstract string) {
	if app.isDeferredService(abstract) && !app.Bound(abstract) {
		app.loadDeferredProvider(abstract)
	}
}

// Determine if the given service is a deferred service.
func (app *Application) isDeferredService(service string) bool {
	_, ok := app.deferredServices[service]
	return ok
}

// Load the provider for a deferred service.
func (app *Application) loadDeferredProvider(service string) {
	if !app.isDeferredService(service) {
		return
	}
	provider := app.deferredServices[service]

	// If the service provider has not already been loaded and registered we can
	// register it with the application and remove the service from this list
	// of deferred services, since it will already be loaded on subsequent.
	if _, ok := app.loadedProviders[app.AbstractToString(provider)]; !ok {
		app.RegisterDeferredProvider(provider, service)
	}
}

// RegisterDeferredProvider Register a deferred provider and service.
func (app *Application) RegisterDeferredProvider(provider support.ServiceProviderInterface, service string) {
	// Once the provider that provides the deferred service has been registered we
	// will remove it from our local list of the deferred services with related
	// providers so that this container does not try to resolve it out again.
	if _, ok := app.deferredServices[service]; ok {
		delete(app.deferredServices, service)
	}

	provider.SetApplication(app)
	app.Register(provider, false)

	if !app.IsBooted() {
		app.Booting(func(application foundation.ApplicationInterface) {
			app.bootProvider(provider)
		})
	}
}

// Booting Register a new boot listener.
func (app *Application) Booting(callback func(app foundation.ApplicationInterface)) {
	app.bootingCallbacks = append(app.bootingCallbacks, callback)
}

// Boot the given service provider.
func (app *Application) bootProvider(provider support.ServiceProviderInterface) {
	provider.CallBootingCallbacks()
	provider.Boot()
	provider.CallBootedCallbacks()
}

// Booted Register a new "booted" listener.
func (app *Application) Booted(callback func(app foundation.ApplicationInterface)) {
	app.bootedCallbacks = append(app.bootedCallbacks, callback)
	if app.IsBooted() {
		callback(app)
	}
}

// LoadDeferredProviders Load and boot all the remaining deferred providers.
func (app *Application) LoadDeferredProviders() {
	// We will simply spin through each of the deferred providers and register each
	// one and boot them if the application has booted. This should make each of
	// the remaining services available to this application for immediate use.
	for service, _ := range app.deferredServices {
		app.loadDeferredProvider(service)
	}

	app.deferredServices = map[string]support.ServiceProviderInterface{}
}
