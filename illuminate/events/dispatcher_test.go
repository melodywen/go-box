package events

import (
	"fmt"
	"github.com/melodywen/go-box/illuminate/contracts/events"
	"github.com/melodywen/go-box/mock"
	container "github.com/melodywen/go-ioc"
	"reflect"
	"testing"
)

func TestDispatcher_Listen(t *testing.T) {
	type fields struct {
		app            interface{}
		listeners      map[string][]events.WrapListenerFun
		wildcards      map[string][]events.WrapListenerFun
		wildcardsCache map[string][]events.WrapListenerFun
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
				events: []interface{}{"ab*c", 123456789, mock.Dog{}},
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
				app:            container.NewContainer(),
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

func TestDispatcher_HasListeners(t *testing.T) {
	type fields struct {
		app            interface{}
		listeners      map[string][]events.WrapListenerFun
		wildcards      map[string][]events.WrapListenerFun
		wildcardsCache map[string][]events.WrapListenerFun
		queueResolver  interface{}
	}
	type args struct {
		event    interface{}
		events   interface{}
		listener events.ListenerFun
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "测试标准值",
			fields: fields{
				listeners:      map[string][]events.WrapListenerFun{},
				wildcards:      map[string][]events.WrapListenerFun{},
				wildcardsCache: map[string][]events.WrapListenerFun{},
			},
			args: args{
				event:  "abc",
				events: "abc",
				listener: func(args ...interface{}) interface{} {
					return 1
				},
			},
			want: true,
		}, {
			name: "测试模糊匹配",
			fields: fields{
				listeners:      map[string][]events.WrapListenerFun{},
				wildcards:      map[string][]events.WrapListenerFun{},
				wildcardsCache: map[string][]events.WrapListenerFun{},
			},
			args: args{
				event:  "abbcc",
				events: "ab*c",
				listener: func(args ...interface{}) interface{} {
					return 1
				},
			},
			want: true,
		}, {
			name: "测试模糊匹配完整匹配",
			fields: fields{
				listeners:      map[string][]events.WrapListenerFun{},
				wildcards:      map[string][]events.WrapListenerFun{},
				wildcardsCache: map[string][]events.WrapListenerFun{},
			},
			args: args{
				event:  "ab*c",
				events: "ab*c",
				listener: func(args ...interface{}) interface{} {
					return 1
				},
			},
			want: true,
		}, {
			name: "测试模糊匹配未匹配到",
			fields: fields{
				listeners:      map[string][]events.WrapListenerFun{},
				wildcards:      map[string][]events.WrapListenerFun{},
				wildcardsCache: map[string][]events.WrapListenerFun{},
			},
			args: args{
				event:  "acbc",
				events: "ab*c",
				listener: func(args ...interface{}) interface{} {
					return 1
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dispatcher := &Dispatcher{
				app:            container.NewContainer(),
				listeners:      tt.fields.listeners,
				wildcards:      tt.fields.wildcards,
				wildcardsCache: tt.fields.wildcardsCache,
				queueResolver:  tt.fields.queueResolver,
			}
			dispatcher.Listen(tt.args.events, tt.args.listener)

			if got := dispatcher.HasListeners(tt.args.event); got != tt.want {
				t.Errorf("HasListeners() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDispatcher_Dispatch(t *testing.T) {

	type args struct {
		event    interface{}
		payload  interface{}
		halt     bool
		events   interface{}
		listener events.ListenerFun
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "测试标准值",
			args: args{
				event:   "abc",
				events:  "abc",
				payload: mock.Dog{Name: "tom", Age: 13},
				listener: func(args ...interface{}) interface{} {
					value := args[0].(mock.Dog)
					value.Age += 1
					return value
				},
			},
			want: []interface{}{mock.Dog{Name: "tom", Age: 14}},
		}, {
			name: "测试标准值并且模糊匹配",
			args: args{
				event:   "melody-wen",
				events:  "me*n",
				payload: mock.Dog{Name: "tom", Age: 13},
				listener: func(args ...interface{}) interface{} {
					value := args[1].([]interface{})
					dog := value[0].(mock.Dog)
					dog.Age += 1
					return dog
				},
			},
			want: []interface{}{mock.Dog{Name: "tom", Age: 14}},
		}, {
			name: "测试标准值并且模糊匹配 --- 命中缓存",
			args: args{
				event:   "melody-wen",
				events:  nil,
				halt:    true,
				payload: mock.Dog{Name: "tom", Age: 13},
			},
			want: []interface{}{mock.Dog{Name: "tom", Age: 14}},
		}, {
			name: "测试标准值并且模糊匹配 - 直接跳过返回nil",
			args: args{
				event:   "melody-we",
				events:  nil,
				halt:    true,
				payload: mock.Dog{Name: "tom", Age: 13},
				listener: func(args ...interface{}) interface{} {
					return nil
				},
			},
			want: []interface{}{nil},
		}, {
			name: "测试标准值并且模糊匹配 - 返回false",
			args: args{
				event:   mock.Cat{},
				events:  mock.Cat{},
				payload: mock.Dog{Name: "tom", Age: 13},
				listener: func(args ...interface{}) interface{} {
					return false
				},
			},
			want: []interface{}{},
		}, {
			name: "测试结构体",
			args: args{
				event:   mock.Dog{Name: "tom2", Age: 20},
				events:  mock.Dog{},
				payload: mock.Dog{Name: "tom", Age: 13},
				listener: func(args ...interface{}) interface{} {
					return args[0]
				},
			},
			want: []interface{}{mock.Dog{Name: "tom2", Age: 20}},
		},
	}
	dispatcher := NewDispatcher(container.NewContainer())
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.args.events != nil {
				dispatcher.Listen(tt.args.events, tt.args.listener)
			}
			var got interface{}
			if tt.args.halt == true {
				got = dispatcher.Until(tt.args.event, tt.args.payload)
				got = []interface{}{got}
			} else {
				got = dispatcher.Dispatch(tt.args.event, tt.args.payload, tt.args.halt)
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasListeners() = %v, want %v", got, tt.want)
			}
		})
	}
}
