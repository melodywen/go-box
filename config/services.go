package config

import (
	"github.com/melodywen/go-box/illuminate/contracts/support"
	"github.com/melodywen/go-box/mock/providers"
)

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
