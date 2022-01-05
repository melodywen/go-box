package support

import (
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
)

// ServiceProvider struct
type ServiceProvider struct {
	App              foundation.ApplicationInterface
	BootingCallbacks []func()
	BootedCallbacks  []func()
}

// NewServiceProvider Create a new service provider instance
func NewServiceProvider(app foundation.ApplicationInterface) *ServiceProvider {
	return &ServiceProvider{App: app}
}

// Boot Bootstrap any application services.
func (provider *ServiceProvider) Boot() {
}

// Register any application services.
func (provider *ServiceProvider) Register() {

}

// Booting Register a booting callback to be run before the "boot" method is called.
func (provider *ServiceProvider) Booting(callback func()) {
	provider.BootingCallbacks = append(provider.BootingCallbacks, callback)
}

// Booted Register a booted callback to be run after the "boot" method is called.
func (provider *ServiceProvider) Booted(callback func()) {
	provider.BootedCallbacks = append(provider.BootedCallbacks, callback)
}

// CallBootingCallbacks Call the registered booting callbacks.
func (provider *ServiceProvider) CallBootingCallbacks() {
	for _, callback := range provider.BootingCallbacks {
		callback()
	}
}

// CallBootedCallbacks Call the registered booted callbacks.
func (provider *ServiceProvider) CallBootedCallbacks() {
	for _, callback := range provider.BootedCallbacks {
		callback()
	}
}

// SetApplication set app Application interface
func (provider *ServiceProvider) SetApplication(app interface{}) {
	if value, ok := app.(foundation.ApplicationInterface); ok {
		provider.App = value
	}
}
