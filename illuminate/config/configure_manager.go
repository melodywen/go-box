package config

import (
	"fmt"
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"path"
	"reflect"
	"regexp"
	"time"
)

// ConfigureManager struct
type ConfigureManager struct {
	app      foundation.ApplicationInterface
	viper    *viper.Viper
	envViper *viper.Viper
}

// NewConfigureManager new an instance
func NewConfigureManager(app foundation.ApplicationInterface) *ConfigureManager {
	manager := &ConfigureManager{
		app:      app,
		viper:    viper.New(),
		envViper: viper.New(),
	}

	manager.viper.SetConfigType("yaml")
	manager.envViper.AutomaticEnv()
	return manager
}

// WriteConfig writer to cache file
func (manager *ConfigureManager) WriteConfig() {
	err := manager.viper.WriteConfigAs(path.Join(manager.app.BootstrapPath(""), "cache", ".config"))
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error write config file: %w \n", err))
	}
}

// parse value if container env
func (manager *ConfigureManager) parseValueIfContainerEnv(value interface{}) interface{} {
	if value == nil {
		return value
	}
	switch reflect.TypeOf(value).Kind() {
	case reflect.Slice:
		v, ok := value.([]interface{})
		if !ok {
			panic("parse value error")
		}
		for key, item := range v {
			switch reflect.TypeOf(item).Kind() {
			case reflect.String:
				v[key] = manager.parseEnvExpression(item)
			case reflect.Map:
				fallthrough
			case reflect.Slice:
				v[key] = manager.parseValueIfContainerEnv(item)
			}
		}
		value = v
	case reflect.Map:
		v, ok := value.(map[string]interface{})
		if !ok {
			panic("parse value error")
		}
		for key, item := range v {
			switch reflect.TypeOf(item).Kind() {
			case reflect.String:
				v[key] = manager.parseEnvExpression(item)
			case reflect.Map:
				fallthrough
			case reflect.Slice:
				v[key] = manager.parseValueIfContainerEnv(item)
			}
		}
		value = v
	}
	return value
}

// pass env expression
func (manager *ConfigureManager) parseEnvExpression(item interface{}) interface{} {
	re := regexp.MustCompile("^\\$\\{(.+?)(\\|\\|.+?)?\\}$")
	matched := re.FindStringSubmatch(item.(string))
	if len(matched) > 0 {
		envName := matched[1]
		envValue := manager.envViper.Get(envName)
		// read config default value
		if envValue == nil && len(matched[2]) != 0 {
			envValue = matched[2][2:]
		}
		return envValue
	}
	return item
}

// Get can retrieve any value given the key to use.
func (manager *ConfigureManager) Get(key string) interface{} {
	value := manager.viper.Get(key)
	value = manager.parseValueIfContainerEnv(value)
	return value
}

// GetBool get bool
func (manager *ConfigureManager) GetBool(key string) bool {
	return cast.ToBool(manager.Get(key))
}

// GetFloat64 config interface GetFloat64 method
func (manager *ConfigureManager) GetFloat64(key string) float64 {
	return cast.ToFloat64(manager.Get(key))
}

// GetInt config interface GetInt method
func (manager *ConfigureManager) GetInt(key string) int {
	return cast.ToInt(manager.Get(key))
}

// GetIntSlice config interface GetIntSlice method
func (manager *ConfigureManager) GetIntSlice(key string) []int {
	return cast.ToIntSlice(manager.Get(key))
}

// GetString config interface GetString method
func (manager *ConfigureManager) GetString(key string) string {
	return cast.ToString(manager.Get(key))
}

// GetStringMap config interface GetStringMap method
func (manager *ConfigureManager) GetStringMap(key string) map[string]interface{} {
	return cast.ToStringMap(manager.Get(key))
}

// GetStringMapString config interface GetStringMapString method
func (manager *ConfigureManager) GetStringMapString(key string) map[string]string {
	return cast.ToStringMapString(manager.Get(key))
}

// GetStringSlice config interface GetStringSlice method
func (manager *ConfigureManager) GetStringSlice(key string) []string {
	return cast.ToStringSlice(manager.Get(key))
}

// GetTime config interface GetTime method
func (manager *ConfigureManager) GetTime(key string) time.Time {
	return cast.ToTime(manager.Get(key))
}

// GetDuration config interface GetDuration method
func (manager *ConfigureManager) GetDuration(key string) time.Duration {
	return cast.ToDuration(manager.Get(key))
}

// IsSet config interface IsSet method
func (manager *ConfigureManager) IsSet(key string) bool {
	return manager.viper.IsSet(key)
}

// AllSettings config interface AllSettings method
func (manager *ConfigureManager) AllSettings() map[string]interface{} {
	panic("implement me")
}

// Set config interface set method
func (manager *ConfigureManager) Set(key string, value interface{}) {
	manager.viper.Set(key, value)
}
