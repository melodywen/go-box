package log

import (
	"github.com/melodywen/go-box/contracts/foundation"
	"github.com/melodywen/go-box/contracts/log"
	"github.com/melodywen/go-box/support"
)

// LoggerServiceProvider struct
type LoggerServiceProvider struct {
	support.ServiceProvider
}

// NewLoggerServiceProvider new log instance
func NewLoggerServiceProvider(app foundation.ApplicationInterface) *LoggerServiceProvider {
	return &LoggerServiceProvider{ServiceProvider: *support.NewServiceProvider(app)}
}

// Register rewrite register
func (provider *LoggerServiceProvider) Register() {

	provider.App.Singleton("log", func(app foundation.ApplicationInterface) log.LoggerInterface {
		return newLogManager(app)
	})

	result := provider.App.Make("log").(log.LoggerInterface)
	result.Info("已经成功注入了log", nil)
}
