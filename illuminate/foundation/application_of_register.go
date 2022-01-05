package foundation

import (
	"github.com/melodywen/go-box/illuminate/contracts/support"
	"reflect"
)

// Register a service provider with the application.
func (app *Application) Register(provider support.ServiceProviderInterface, force bool) support.ServiceProviderInterface {
	if registered := app.GetProvider(provider); registered != nil && !force {
		return registered
	}

	provider.SetApplication(app)
	provider.Register()

	// If there are bindings / singletons set as properties on the provider we
	// will spin through them and register them with the application, which
	// serves as a convenience layer while registering a lot of bindings.

	// todo bindings
	// todo singletons

	app.markAsRegistered(provider)

	// If the application has already booted, we will call this boot method on
	// the provider class so it has an opportunity to do its boot logic and
	// will be ready for any usage by this developer's application logic.
	if app.IsBooted() {
		app.bootProvider(provider)
	}
	return provider
}

// IsBooted Determine if the application has booted.
func (app *Application) IsBooted() bool {
	return app.booted
}

// GetProvider Get the registered service provider instance if it exists.
func (app *Application) GetProvider(provider support.ServiceProviderInterface) support.ServiceProviderInterface {
	result := app.GetProviders(provider)
	if len(result) == 0 {
		return nil
	}
	return result[0]
}

// GetProviders Get the registered service provider instances if any exist.
func (app *Application) GetProviders(provider support.ServiceProviderInterface) []support.ServiceProviderInterface {
	obj := make([]support.ServiceProviderInterface, 0)
	for _, serviceProvider := range app.serviceProviders {
		if reflect.TypeOf(serviceProvider) == reflect.TypeOf(provider) {
			obj = append(obj, serviceProvider)
		}
	}
	return obj
}

// Mark the given provider as registered.
func (app *Application) markAsRegistered(provider support.ServiceProviderInterface) {
	app.serviceProviders = append(app.serviceProviders, provider)

	app.loadedProviders[app.AbstractToString(provider)] = true
}
