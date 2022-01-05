package bootstrap

import (
	"fmt"
	"github.com/melodywen/go-box/illuminate/config"
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	"github.com/spf13/viper"
	"os"
	"path"
)

// LoadConfiguration load configure struct
type LoadConfiguration struct {
	app              foundation.ApplicationInterface
	configureManager *config.ConfigureManager
}

// NewLoadConfiguration  new an config  bootstrap instance
func NewLoadConfiguration() *LoadConfiguration {
	return &LoadConfiguration{}
}

// Bootstrap Register services.
func (bootstrap *LoadConfiguration) Bootstrap(app foundation.ApplicationInterface) {
	bootstrap.app = app
	bootstrap.configureManager = config.NewConfigureManager(app)

	// Next we will spin through all the configuration files in the configuration
	// directory and load each one into the repository. This will make all of the
	// options available to the developer for use in various parts of this app.
	bootstrap.app.Instance("config", bootstrap.configureManager)

	bootstrap.loadConfigurationFiles(bootstrap.app, bootstrap.configureManager)

}

//loadConfigurationFiles load the configuration items from all the files.
func (bootstrap *LoadConfiguration) loadConfigurationFiles(app foundation.ApplicationInterface, manager *config.ConfigureManager) {
	files := bootstrap.getConfigurationFiles(app)
	configPath := app.ConfigPath("")
	allSuffix := map[string]bool{
		"json": true, "toml": true, "yaml": true,
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := path.Base(file.Name())
		fileSuffix := path.Ext(filename)
		if fileSuffix == "" {
			continue
		}
		if _, ok := allSuffix[fileSuffix[1:]]; !ok {
			continue
		}
		filePrefix := filename[0 : len(filename)-len(fileSuffix)]

		config := viper.New()
		config.SetConfigName(filePrefix)     //name of config file (without extension)
		config.SetConfigType(fileSuffix[1:]) // REQUIRED if the config file does not have the extension in the name
		config.AddConfigPath(configPath)     // path to look for the config file in
		err := config.ReadInConfig()         // Find and read the config file
		if err != nil {                      // Handle errors reading the config file
			panic(fmt.Errorf("Fatal error config file: %w \n", err))
		}
		manager.Set(filePrefix, config.AllSettings())
	}
	manager.WriteConfig()
}

//getConfigurationFiles Get all the configuration files for the application.
func (bootstrap *LoadConfiguration) getConfigurationFiles(app foundation.ApplicationInterface) []os.DirEntry {
	configPath := app.ConfigPath("")
	dir, err := os.ReadDir(configPath)
	if err != nil {
		panic(err)
	}
	return dir
}
