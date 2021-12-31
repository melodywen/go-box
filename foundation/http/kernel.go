package http

import (
	events2 "github.com/melodywen/go-box/contracts/events"
	"github.com/melodywen/go-box/contracts/foundation"
	log2 "github.com/melodywen/go-box/contracts/log"
	"github.com/melodywen/go-box/foundation/bootstrap"
	"github.com/sirupsen/logrus"
)

type Kernel struct {
	app           foundation.ApplicationInterface
	bootstrappers []foundation.BootstrapInterface
}

var bootstraps []foundation.BootstrapInterface

// init global variable
func init() {
	bootstraps = []foundation.BootstrapInterface{
		bootstrap.NewFoundationProvider(),
		bootstrap.NewLoadConfiguration(),
		bootstrap.NewBootProviders(),
	}
}

// NewKernel new a http kernel
func NewKernel(application foundation.ApplicationInterface) *Kernel {
	kernel := &Kernel{
		app: application,
	}
	kernel.bootstrappers = append(kernel.bootstrappers, bootstraps...)
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

// Bootstrap the application for HTTP requests.
func (k *Kernel) Bootstrap() {
	k.bootstrapListen()
	if !k.app.HasBeenBootstrapped() {
		k.app.BootstrapWith(k.getBootstrappers())
	}
}

// bootstrapListen add bootstrap to listen fun
func (k *Kernel) bootstrapListen() {
	var dispatcher events2.DispatcherInterface
	dispatcher = k.app.Make("events").(events2.DispatcherInterface)
	var log log2.LoggerInterface
	log = k.app.Make("log").(log2.LoggerInterface)

	dispatcher.Listen("bootstrapping:*", func(args ...interface{}) interface{} {
		log.Info(args[0].(string), nil)
		return nil
	})
	dispatcher.Listen("bootstrapped:*", func(args ...interface{}) interface{} {
		log.Info(args[0].(string), nil)
		return nil
	})
}

// Handle init kernel
func (k *Kernel) Handle() {
	k.Bootstrap()
}

func (k *Kernel) Terminate() {
	panic("implement me")
}

func (k *Kernel) GetApplication() {
	panic("implement me")
}
