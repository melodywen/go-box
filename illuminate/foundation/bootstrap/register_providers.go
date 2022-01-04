package bootstrap

import (
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
)

type RegisterProviders struct {
}

func NewRegisterProviders() *RegisterProviders {
	return &RegisterProviders{}
}


func (r RegisterProviders) Bootstrap(app foundation.ApplicationInterface) {
	app.RegisterConfiguredProviders()
}
