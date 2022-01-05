package cache

import (
	"fmt"
	"github.com/melodywen/go-box/illuminate/support"
	"reflect"
)

//CacheServiceProvider d
type CacheServiceProvider struct {
	support.ServiceProvider
}

func NewCacheServiceProvider() *CacheServiceProvider {
	return &CacheServiceProvider{}
}

// Boot Bootstrap any application services.
func (provider *CacheServiceProvider) Boot() {
}

// Register any application services.
func (provider *CacheServiceProvider) Register() {
	fmt.Println(reflect.TypeOf(provider.App))
	fmt.Println("注册cache", provider)
}
