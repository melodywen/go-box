package foundation

import (
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	"github.com/melodywen/go-box/illuminate/contracts/support"
)

// ProviderRepository provider repository struct
type ProviderRepository struct {
	app foundation.ApplicationInterface
}

// NewProviderRepository new an instance
func NewProviderRepository(app foundation.ApplicationInterface) *ProviderRepository {
	return &ProviderRepository{app: app}
}

// Load Register the application service providers.
func (repository *ProviderRepository) Load() {
	eagerServices := repository.app.Make("eager-services").([]support.ServiceProviderInterface)

	// Next, we will register events to load the providers for each of the events
	// that it has requested. This allows the service provider to defer itself
	// while still getting automatically loaded when a certain event occurs.
	// todo license event to register

	// We will go ahead and register all of the eagerly loaded providers with the
	// application so their services can be registered with the application as
	// a provided service. Then we will set the deferred service list on it.
	for _, service := range eagerServices {
		repository.app.Register(service, false)
	}

	deferServices := repository.app.Make("defer-services").(map[string]support.ServiceProviderInterface)
	repository.app.AddDeferredServices(deferServices)
}
