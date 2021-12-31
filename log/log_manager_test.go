package log

import (
	"github.com/melodywen/go-box/contracts/foundation"
	"testing"
)

func Test_newLogManager(t *testing.T) {
	type args struct {
		app foundation.ApplicationInterface
	}
	tests := []struct {
		name string
		args args
		want *LoggerManager
	}{
		{
			name: "测试默认驱动",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			log := NewLoggerManager(nil)

			log.Trace("a", nil)
			log.Debug("a", nil)
			log.Info("a", nil)
			log.Warn("a", nil)
			log.Error("a", nil)
			defer func() {
				recover()
			}()
			log.Panic("a", nil)

		})
	}
}
