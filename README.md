# go-box

[![Build Status](https://github.com/gin-gonic/gin/workflows/Run%20Tests/badge.svg?branch=master)]()
[![Go Report Card](https://camo.githubusercontent.com/f05145ad1c938e873697d2b624764921913522654e41fb7c68ba7918967a846b/68747470733a2f2f676f7265706f7274636172642e636f6d2f62616467652f6769746875622e636f6d2f676f2d676f726d2f676f726d)]()
[![GoDoc](https://pkg.go.dev/badge/github.com/gin-gonic/gin?status.svg)]()
[![license](https://camo.githubusercontent.com/992daabc2aa4463339825f8333233ba330dd08c57068f6faf4bb598ab5a3df2e/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f6c6963656e73652d4d49542d627269676874677265656e2e737667)]()

## Overview
This project builds the necessary components such as configuration files, service providers, middleware, Facades, and so on. The specific implementation is to refer to laravel framework implementation;

### done:
- Configuration file
- Service provider
- Event dispatcher
- Overall improvement of the framework of application

### todo:
- Exception handling
- facades
- middleware

## example
config/services.go
```go
// EagerServices eager services
var EagerServices []support.ServiceProviderInterface

// DeferServices defer services
var DeferServices map[string]support.ServiceProviderInterface

func init() {
	EagerServices = []support.ServiceProviderInterface{
		providers.NewSchoolServiceProvider(),
	}
	DeferServices = map[string]support.ServiceProviderInterface{
		"teacher": providers.NewTeacherServiceProvider(),
		"student": providers.NewStudentServiceProvider(),
	}
}
```

app.go
```golang
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
```

main.go
```golang
func main() {
	var httpKernel http.KernelInterface
	var ok bool
	k := App.Make(&httpKernel)

	if httpKernel, ok = k.(http.KernelInterface); !ok {
		logrus.Panicln("获取 http kernel 失败")
	}

	httpKernel.Handle()

	App.Make("school")
	fmt.Println(httpKernel)
}
```