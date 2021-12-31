package foundation

import (
	"github.com/melodywen/go-box/contracts/support"
	"github.com/melodywen/go-ioc/contracts"
)

// BootstrapInterface bootstrap contract
type BootstrapInterface interface {
	// Bootstrap Register services.
	Bootstrap(app ApplicationInterface)
}

// ApplicationInterface application contract
type ApplicationInterface interface {
	contracts.ContainerContract
	//Version Get the version number of the application.
	Version() string
	//BasePath Get the base path of the Laravel installation.
	BasePath()
	// BootstrapPath Get the path to the bootstrap directory.
	BootstrapPath()
	// ConfigPath Get the path to the application configuration files.
	ConfigPath()
	// DatabasePath Get the path to the database directory.
	DatabasePath()
	// ResourcePath Get the path to the resources directory.
	ResourcePath()
	// StoragePath Get the path to the storage directory.
	StoragePath()
	// Environment Get or check the current application environment.
	Environment()
	// RunningInConsole Determine if the application is running in the console.
	RunningInConsole()
	// RunningUnitTests Determine if the application is running unit tests.
	RunningUnitTests()
	// IsDownForMaintenance Determine if the application is currently down for maintenance.
	IsDownForMaintenance()
	// RegisterConfiguredProviders Register all of the configured providers.
	RegisterConfiguredProviders()
	//Register a service provider with the application.
	Register(provider support.ServiceProviderInterface, force bool) support.ServiceProviderInterface
	// RegisterDeferredProvider Register a deferred provider and service.
	RegisterDeferredProvider()
	// ResolveProvider Resolve a service provider instance from the class name.
	ResolveProvider()
	// Boot the application's service providers.
	Boot()
	// Booting Register a new boot listener.
	Booting()
	// Booted Register a new "booted" listener.
	Booted()
	// BootstrapWith Run the given array of bootstrap classes.
	BootstrapWith(bootstrappers []BootstrapInterface)
	// GetLocale Get the current application locale.
	GetLocale()
	// GetNamespace Get the application namespace.
	GetNamespace()
	// GetProviders Get the registered service provider instances if any exist.
	GetProviders(provider support.ServiceProviderInterface) []support.ServiceProviderInterface
	// HasBeenBootstrapped Determine if the application has been bootstrapped before.
	HasBeenBootstrapped() bool
	// LoadDeferredProviders Load and boot all of the remaining deferred providers.
	LoadDeferredProviders()
	// SetLocale Set the current application locale.
	SetLocale()
	// ShouldSkipMiddleware Determine if middleware has been disabled for the application.
	ShouldSkipMiddleware()
	// Terminate the application.
	Terminate()
}
