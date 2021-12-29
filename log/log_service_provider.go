package log

import (
	"fmt"
	"github.com/melodywen/go-box/contracts/foundation"
	"github.com/melodywen/go-box/support"
)

type LogServiceProvider struct {
	support.ServiceProvider
}

func NewLogServiceProvider(app foundation.ApplicationInterface) *LogServiceProvider {
	return &LogServiceProvider{ServiceProvider: *support.NewServiceProvider(app)}
}

func (provider *LogServiceProvider) Boot() {

}

func (provider *LogServiceProvider) Register() {
	provider.App.Singleton("log", func(app foundation.ApplicationInterface) {
		fmt.Println("这是一个 log 的具体实现")
	})
}
