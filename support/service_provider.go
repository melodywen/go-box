package support

import "github.com/melodywen/go-box/contracts/foundation"

type ServiceProvider struct {
	App              foundation.ApplicationInterface
	BootingCallbacks []func()
	BootedCallbacks  []func()
}

func NewServiceProvider(app foundation.ApplicationInterface) *ServiceProvider {
	return &ServiceProvider{App: app}
}

func (provider *ServiceProvider) Register() {

}
func (provider *ServiceProvider) Booting() {
	panic("implement me")
}
func (provider *ServiceProvider) Booted() {
	panic("implement me")
}
func (provider *ServiceProvider) CallBootingCallbacks() {
	panic("implement me")
}
func (provider *ServiceProvider) CallBootedCallbacks() {
	panic("implement me")
}
