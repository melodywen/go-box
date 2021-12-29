package http

import (
	"github.com/melodywen/go-box/contracts/foundation"
	"github.com/sirupsen/logrus"
)

type Kernel struct {
	app foundation.ApplicationInterface

	bootstrappers []foundation.BootstrapInterface
}

func NewKernel(application foundation.ApplicationInterface) *Kernel {
	kernel := &Kernel{
		app: application,
	}
	kernel.SyncMiddlewareToRouter()
	return kernel
}

// SyncMiddlewareToRouter 加载中间件
func (k *Kernel) SyncMiddlewareToRouter() {
	logrus.Warnln("todo: 待实现 SyncMiddlewareToRouter")
}

//getBootstrappers
func (k *Kernel) getBootstrappers() []foundation.BootstrapInterface {
	return k.bootstrappers
}

func (k *Kernel) Bootstrap() {
	if !k.app.HasBeenBootstrapped() {
		k.app.BootstrapWith(k.getBootstrappers())
	}
}

func (k *Kernel) Handle() {
	panic("implement me")
}

func (k *Kernel) Terminate() {
	panic("implement me")
}

func (k *Kernel) GetApplication() {
	panic("implement me")
}
