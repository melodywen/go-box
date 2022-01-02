package bootstrap

import (
	"github.com/melodywen/go-box/contracts/foundation"
)

type LoadConfiguration struct {
	app              foundation.ApplicationInterface
	configureManager *ConfigureManager
}

func NewLoadConfiguration() *LoadConfiguration {
	return &LoadConfiguration{}
}

func (bootstrap *LoadConfiguration) Bootstrap(app foundation.ApplicationInterface) {
	bootstrap.app = app
	bootstrap.configureManager = NewConfigureManager()

	// Next we will spin through all the configuration files in the configuration
	// directory and load each one into the repository. This will make all of the
	// options available to the developer for use in various parts of this app.
	bootstrap.app.Instance("config", bootstrap.configureManager)

	bootstrap.loadConfigurationFiles(bootstrap.app, bootstrap.configureManager)

}

//loadConfigurationFiles load the configuration items from all the files.
func (bootstrap *LoadConfiguration) loadConfigurationFiles(app foundation.ApplicationInterface, manager *ConfigureManager) {
	bootstrap.getConfigurationFiles(app)
}

//getConfigurationFiles Get all the configuration files for the application.
func (bootstrap *LoadConfiguration) getConfigurationFiles(app foundation.ApplicationInterface) {


}
