package bootstrap

import (
	"github.com/melodywen/go-box/contracts/foundation"
)

type FoundationProvider struct {
	app foundation.ApplicationInterface
}


func NewFoundationProvider() *FoundationProvider {
	return &FoundationProvider{}
}

// Bootstrap Register foundation provider services.
func (foundation *FoundationProvider) Bootstrap(app foundation.ApplicationInterface) {
}
