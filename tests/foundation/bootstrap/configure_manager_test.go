package bootstrap

import (
	"fmt"
	"github.com/melodywen/go-box/config"
	config2 "github.com/melodywen/go-box/illuminate/config"
	"github.com/melodywen/go-box/illuminate/contracts/http"
	foundation2 "github.com/melodywen/go-box/illuminate/foundation"
	http2 "github.com/melodywen/go-box/illuminate/foundation/http"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"testing"
)

func TestConfigureManager_WriteConfig(t *testing.T) {

	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "基础测试",
		},
	}
	dir, _ := os.Getwd()
	dir = path.Join(dir, "../../../")
	app := foundation2.NewApplication(dir)

	app.BootstrapOpenListen()
	app.Instance("eager-services", config.EagerServices)
	app.Instance("defer-services", config.DeferServices)
	var httpKernel http.KernelInterface
	app.Singleton(&httpKernel, http2.NewKernel)
	var ok bool
	k := app.Make(&httpKernel)

	if httpKernel, ok = k.(http.KernelInterface); !ok {
		logrus.Panicln("获取 http kernel 失败")
	}
	httpKernel.Handle()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := app.Make("config")

			var config *config2.ConfigureManager
			if config, ok = v.(*config2.ConfigureManager); !ok {
				logrus.Panicln("获取 config 失败")
			}
			value := config.Get("test_struct.driver")
			fmt.Println(value)
			value = config.Get("test_struct.driver1")
			fmt.Println(value)
			value = config.Get("test_struct.driver2")
			fmt.Println(value)
			value = config.Get("test_struct.driver3")
			fmt.Println(value)
			value = config.GetBool("test_struct.is_active")
			fmt.Println(value)
			value = config.GetFloat64("test_struct.time")
			fmt.Println(value)
		})
	}
}
