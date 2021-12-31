package bootstrap

import "github.com/melodywen/go-box/contracts/foundation"

type LoadConfiguration struct {
	app foundation.ApplicationInterface
}

func NewLoadConfiguration() *LoadConfiguration {
	return &LoadConfiguration{}
}

func (config *LoadConfiguration) Bootstrap(app foundation.ApplicationInterface) {

}
