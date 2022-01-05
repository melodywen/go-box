package bootstrap

import (
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
)

// RegisterProviders register providers struct
type RegisterProviders struct {
}

// NewRegisterProviders  new an instance
func NewRegisterProviders() *RegisterProviders {
	return &RegisterProviders{}
}

// Bootstrap provider bootstrap
func (r RegisterProviders) Bootstrap(app foundation.ApplicationInterface) {
	app.RegisterConfiguredProviders()
}
