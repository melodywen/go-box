package support

// ServiceProviderInterface base service provider interface
type ServiceProviderInterface interface {
	// Register any application services.
	Register()
	// Boot Bootstrap any application services.
	Boot()
	// Booting Register a booting callback to be run before the "boot" method is called.
	Booting(callback func())
	// Booted Register a booted callback to be run after the "boot" method is called.
	Booted(callback func())
	// CallBootingCallbacks Call the registered booting callbacks.
	CallBootingCallbacks()
	// CallBootedCallbacks Call the registered booted callbacks.
	CallBootedCallbacks()
}
