package bootstrap

import (
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
)

// BootProviders boot provider struct
type BootProviders struct {
	app foundation.ApplicationInterface
}

// NewBootProviders new boot provider instance
func NewBootProviders() *BootProviders {
	return &BootProviders{}
}

// Bootstrap the given application.
func (boot *BootProviders) Bootstrap(app foundation.ApplicationInterface) {
	app.Boot()
}
