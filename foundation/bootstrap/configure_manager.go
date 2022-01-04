package bootstrap

import (
	"fmt"
	"github.com/melodywen/go-box/contracts/foundation"
	"github.com/spf13/viper"
	"path"
	"time"
)

type ConfigureManager struct {
	app   foundation.ApplicationInterface
	viper *viper.Viper
}

func NewConfigureManager(app foundation.ApplicationInterface) *ConfigureManager {
	manager := &ConfigureManager{
		app:   app,
		viper: viper.New(),
	}

	manager.viper.SetConfigType("yaml")
	return manager
}

func (manager *ConfigureManager) WriteConfig() {
	err := manager.viper.WriteConfigAs(path.Join(manager.app.BootstrapPath(""), "cache", ".config"))
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error write config file: %w \n", err))
	}
}

func (manager *ConfigureManager) Get(key string) interface{} {
	panic("implement me")
}

func (manager *ConfigureManager) GetBool(key string) bool {
	panic("implement me")
}

func (manager *ConfigureManager) GetFloat64(key string) float64 {
	panic("implement me")
}

func (manager *ConfigureManager) GetInt(key string) int {
	panic("implement me")
}

func (manager *ConfigureManager) GetIntSlice(key string) []int {
	panic("implement me")
}

func (manager *ConfigureManager) GetString(key string) string {
	panic("implement me")
}

func (manager *ConfigureManager) GetStringMap(key string) map[string]interface{} {
	panic("implement me")
}

func (manager *ConfigureManager) GetStringMapString(key string) map[string]string {
	panic("implement me")
}

func (manager *ConfigureManager) GetStringSlice(key string) []string {
	panic("implement me")
}

func (manager *ConfigureManager) GetTime(key string) time.Time {
	panic("implement me")
}

func (manager *ConfigureManager) GetDuration(key string) time.Duration {
	panic("implement me")
}

func (manager *ConfigureManager) IsSet(key string) bool {
	panic("implement me")
}

func (manager *ConfigureManager) AllSettings() map[string]interface{} {
	panic("implement me")
}

func (manager *ConfigureManager) Set(key string, value interface{}) {
	manager.viper.Set(key, value)
}
