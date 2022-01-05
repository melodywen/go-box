package main

import (
	"github.com/melodywen/go-box/config"
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	"github.com/melodywen/go-box/illuminate/contracts/http"
	foundation2 "github.com/melodywen/go-box/illuminate/foundation"
	http2 "github.com/melodywen/go-box/illuminate/foundation/http"
	"os"
)

// App construction app instance
var App foundation.ApplicationInterface

// init App
func init() {
	dir, _ := os.Getwd()

	app := foundation2.NewApplication(dir)
	app.Instance("eager-services", config.EagerServices)
	app.Instance("defer-services", config.DeferServices)

	App = app
	app.BootstrapOpenListen()
	var httpKernel http.KernelInterface
	App.Singleton(&httpKernel, http2.NewKernel)
}
