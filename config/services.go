package config

import "github.com/melodywen/go-box/contracts/support"

var EagerServices []support.ServiceProviderInterface

var DeferServices map[string]support.ServiceProviderInterface

func init() {
	EagerServices = []support.ServiceProviderInterface{}
	DeferServices = map[string]support.ServiceProviderInterface{}
}
