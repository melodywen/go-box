package foundation

import (
	"github.com/melodywen/go-box/contracts/foundation"
	log2 "github.com/melodywen/go-box/contracts/log"
	"github.com/melodywen/go-box/contracts/support"
	"github.com/melodywen/go-box/events"
	"github.com/melodywen/go-box/log"
	container "github.com/melodywen/go-ioc"
	"github.com/sirupsen/logrus"
	"reflect"
)

// Application app struct
type Application struct {
	container.Container
	hasBeenBootstrapped bool                               //Indicates if the application has been bootstrapped before.
	serviceProviders    []support.ServiceProviderInterface // All of the registered service providers.
	loadedProviders     map[string]bool                    //The names of the loaded service providers.
	Log                 log2.LoggerInterface
	booted              bool
}

// NewApplication Create a new Illuminate application instance.
func NewApplication() *Application {
	app := &Application{
		Container:        *container.NewContainer(),
		serviceProviders: []support.ServiceProviderInterface{},
		loadedProviders:  map[string]bool{},
	}

	app.registerBaseBindings()
	app.registerBaseServiceProviders()
	app.registerCoreContainerAliases()
	return app
}

func (app *Application) registerBaseBindings() {
	var App foundation.ApplicationInterface
	app.Instance(&App, app)
	app.Instance(&app, app)
}

// Register all of the base service providers.
func (app *Application) registerBaseServiceProviders() {
	app.Register(events.NewEventServiceProvider(app), false)
	app.Register(log.NewLoggerServiceProvider(app), false)
	logrus.Warnln("todo: 待实现 RoutingServiceProvider")
}

// Register the core class aliases in the container.
func (app *Application) registerCoreContainerAliases() {
	var loggerInterface log2.LoggerInterface
	aliases := map[string][]interface{}{
		"app":    []interface{}{app, container.Container{}},
		"events": []interface{}{},
		"log":    []interface{}{&loggerInterface},
	}
	for key, aliases := range aliases {
		for _, alias := range aliases {
			app.Alias(key, alias)
		}
	}
}

// Register a service provider with the application.
func (app *Application) Register(provider support.ServiceProviderInterface, force bool) support.ServiceProviderInterface {
	if registered := app.GetProvider(provider); registered != nil && !force {
		return registered
	}

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
	if app.isBooted() {

	}

	return provider
}

// Mark the given provider as registered.
func (app *Application) markAsRegistered(provider support.ServiceProviderInterface) {
	app.serviceProviders = append(app.serviceProviders, provider)

	app.loadedProviders[app.AbstractToString(provider)] = true
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

// Determine if the application has booted.
func (app *Application) isBooted() bool {
	return app.booted
}

// Boot the given service provider.
func (app *Application) bootProvider(provider support.ServiceProviderInterface) {
	provider.CallBootingCallbacks()
	provider.Boot()
	provider.CallBootedCallbacks()
}

func (app *Application) HasBeenBootstrapped() bool {
	return app.hasBeenBootstrapped
}

func (app *Application) BootstrapWith(bootstrappers []foundation.BootstrapInterface) {
	app.hasBeenBootstrapped = true
	for _, bootstrapper := range bootstrappers {
		bootstrapper.Bootstrap(app)
	}
}
