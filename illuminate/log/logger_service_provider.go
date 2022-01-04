package log

import (
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	"github.com/melodywen/go-box/illuminate/contracts/log"
	"github.com/melodywen/go-box/illuminate/support"
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
		return NewLoggerManager(app)
	})
}
