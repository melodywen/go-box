package events

import (
	"fmt"
	"github.com/melodywen/go-box/contracts/events"
	"github.com/melodywen/go-box/mock"
	container "github.com/melodywen/go-ioc"
	"testing"
)

func TestDispatcher_Listen(t *testing.T) {
	type fields struct {
		app interface{}
		listeners      map[string][]events.WrapListenerFun
		wildcards      map[string][]events.WrapListenerFun
		wildcardsCache []string
		queueResolver  interface{}
	}
	type args struct {
		events   interface{}
		listener events.ListenerFun
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "第一个验证:普通的注入",
			fields: fields{
				listeners: map[string][]events.WrapListenerFun{},
				wildcards: map[string][]events.WrapListenerFun{},
			},
			args: args{
				events: "abc",
				listener: func(args ...interface{}) interface{} {
					return 1
				},
			},
		}, {
			name: "第二个注入模糊",
			fields: fields{
				listeners: map[string][]events.WrapListenerFun{},
				wildcards: map[string][]events.WrapListenerFun{},
			},
			args: args{
				events: []interface{}{"ab*c",123456789, mock.Dog{}},
				listener: func(args ...interface{}) interface{} {
					fmt.Println("i am is listener")
					return 1
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Dispatcher{
				app: container.NewContainer(),
				listeners:      tt.fields.listeners,
				wildcards:      tt.fields.wildcards,
				wildcardsCache: tt.fields.wildcardsCache,
				queueResolver:  tt.fields.queueResolver,
			}
			d.Listen(tt.args.events, tt.args.listener)

			fmt.Println(d)
		})
	}
}
