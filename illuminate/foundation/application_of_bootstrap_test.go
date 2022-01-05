package foundation

import (
	"fmt"
	"github.com/melodywen/go-box/config"
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
	"github.com/melodywen/go-box/illuminate/contracts/http"
	http2 "github.com/melodywen/go-box/illuminate/foundation/http"
	"github.com/melodywen/go-box/mock/providers"
	"github.com/sirupsen/logrus"
	"os"
	"path"
	"testing"
)

func TestApplication_BootstrapOpenListen(t *testing.T) {

	tests := []struct {
		name string
	}{
		{name: "测试一轮"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dir, _ := os.Getwd()
			dir = path.Join(dir, "../../")
			app := NewApplication(dir)
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
			app.Register(providers.NewSchoolServiceProvider(), true)
			app.Booted(func(app foundation.ApplicationInterface) {
				fmt.Println("booted callback")
			})
			app.Make("school")
			app.LoadDeferredProviders()
		})
	}
}
