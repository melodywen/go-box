package foundation

import (
	"fmt"
	"github.com/melodywen/go-box/contracts/support"
	"github.com/melodywen/go-box/events"
	"github.com/melodywen/go-box/log"
	"reflect"
	"testing"
)

func TestApplication_GetProviders(t *testing.T) {
	app := NewApplication()

	type fields struct {
		serviceProviders []support.ServiceProviderInterface
	}
	type args struct {
		provider support.ServiceProviderInterface
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []support.ServiceProviderInterface
	}{
		{
			name: "第一组测试",
			fields: fields{serviceProviders: []support.ServiceProviderInterface{
				events.NewEventServiceProvider(app),
			}},
			args: args{provider: events.NewEventServiceProvider(app)},
			want: []support.ServiceProviderInterface{events.NewEventServiceProvider(app)},
		}, {
			name: "第二组测试",
			fields: fields{serviceProviders: []support.ServiceProviderInterface{
				events.NewEventServiceProvider(app),
			}},
			args: args{provider: log.NewLogServiceProvider(app)},
			want: []support.ServiceProviderInterface{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			application := &Application{
				serviceProviders: tt.fields.serviceProviders,
			}
			got := application.GetProviders(tt.args.provider)
			fmt.Println(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetProviders() = %v, want %v", got, tt.want)
			}
		})
	}
}