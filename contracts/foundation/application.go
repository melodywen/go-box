package foundation

import (
	"github.com/melodywen/go-box/contracts/support"
	"github.com/melodywen/go-ioc/contracts"
)

type BootstrapInterface interface {
	Bootstrap(app ApplicationInterface)
}

type ApplicationInterface interface {
	contracts.ContainerContract

	//Version Get the version number of the application.
	//Version()string
	// BasePath Get the base path of the Laravel installation.
	//BasePath()

	//Register a service provider with the application.
	Register(provider support.ServiceProviderInterface, force bool) support.ServiceProviderInterface

	// GetProviders Get the registered service provider instances if any exist.
	GetProviders(provider support.ServiceProviderInterface) []support.ServiceProviderInterface

	// HasBeenBootstrapped Determine if the application has been bootstrapped before.
	HasBeenBootstrapped() bool

	// BootstrapWith Run the given array of bootstrap classes.
	BootstrapWith(bootstrappers []BootstrapInterface)
}
