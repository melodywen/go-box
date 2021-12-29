package main

import "C"
import (
	"github.com/melodywen/go-box/contracts/foundation"
	"github.com/melodywen/go-box/contracts/http"
	foundation2 "github.com/melodywen/go-box/foundation"
	http2 "github.com/melodywen/go-box/foundation/http"
)

// App construction app instance
var App foundation.ApplicationInterface

// init App
func init() {
	App = foundation2.NewApplication()

	var httpKernel http.KernelInterface
	App.Singleton(&httpKernel, http2.NewKernel)
}
