package log

import (
	"fmt"
	"github.com/melodywen/go-box/contracts/foundation"
	"github.com/melodywen/go-box/support"
)

// LogServiceProvider struct
type LogServiceProvider struct {
	support.ServiceProvider
}

// NewLogServiceProvider new log instance
func NewLogServiceProvider(app foundation.ApplicationInterface) *LogServiceProvider {
	return &LogServiceProvider{ServiceProvider: *support.NewServiceProvider(app)}
}

// Register rewrite register
func (provider *LogServiceProvider) Register() {
	provider.App.Singleton("log", func(app foundation.ApplicationInterface) {
		fmt.Println("这是一个 log 的具体实现")
	})
}
