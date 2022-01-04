package foundation

import (
	events2 "github.com/melodywen/go-box/illuminate/contracts/events"
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	log2 "github.com/melodywen/go-box/illuminate/contracts/log"
	"github.com/melodywen/go-box/illuminate/contracts/support"
	events3 "github.com/melodywen/go-box/illuminate/events"
	"github.com/melodywen/go-box/illuminate/log"
	container "github.com/melodywen/go-ioc"
	"github.com/sirupsen/logrus"
)

var version string = "1.0.0"

// Application app struct
type Application struct {
	container.Container
	hasBeenBootstrapped bool                               //Indicates if the application has been bootstrapped before.
	serviceProviders    []support.ServiceProviderInterface // All of the registered service providers.
	loadedProviders     map[string]bool                    //The names of the loaded service providers.

	basePath string // base path for the application.

	Log    log2.LoggerInterface
	booted bool
}

// NewApplication Create a new Illuminate application instance.
func NewApplication(basePath string) *Application {
	app := &Application{
		Container:        *container.NewContainer(),
		serviceProviders: []support.ServiceProviderInterface{},
		loadedProviders:  map[string]bool{},
	}

	if basePath != "" {
		app.setBasePath(basePath)
	}

	app.registerBaseBindings()
	app.registerBaseServiceProviders()
	app.registerCoreContainerAliases()
	return app
}

// Register the basic bindings into the container.
func (app *Application) registerBaseBindings() {
	var App foundation.ApplicationInterface
	app.Instance(&App, app)
	app.Instance(&app, app)
}

// Register all of the base service providers.
func (app *Application) registerBaseServiceProviders() {
	app.Register(events3.NewEventServiceProvider(app), false)
	app.Register(log.NewLoggerServiceProvider(app), false)
	logrus.Warnln("todo: 待实现 RoutingServiceProvider")
}

// Register the core class aliases in the container.
func (app *Application) registerCoreContainerAliases() {
	var loggerInterface log2.LoggerInterface
	var dispatcherInterface events2.DispatcherInterface
	aliases := map[string][]interface{}{
		"app":    []interface{}{app, container.Container{}},
		"events": []interface{}{&dispatcherInterface, events3.Dispatcher{}, &events3.Dispatcher{}},
		"log":    []interface{}{&loggerInterface},
	}
	for key, aliases := range aliases {
		for _, alias := range aliases {
			app.Alias(key, alias)
		}
	}
}

func (app *Application) Version() string {
	return version
}