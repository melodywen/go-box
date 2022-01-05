package bootstrap

import (
	"github.com/melodywen/go-box/illuminate/contracts/foundation"
)

// FoundationProvider foundation provider struct
type FoundationProvider struct {
	app foundation.ApplicationInterface
}

// NewFoundationProvider an instance of foundation provider
func NewFoundationProvider() *FoundationProvider {
	return &FoundationProvider{}
}

// Bootstrap Register foundation provider services.
func (foundation *FoundationProvider) Bootstrap(app foundation.ApplicationInterface) {
}
