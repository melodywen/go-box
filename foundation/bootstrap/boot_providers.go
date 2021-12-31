package bootstrap

import (
	"fmt"
	"github.com/melodywen/go-box/contracts/foundation"
)

type BootProviders struct {
	app foundation.ApplicationInterface
}

func NewBootProviders() *BootProviders {
	return &BootProviders{}
}

func (boot *BootProviders) Bootstrap(app foundation.ApplicationInterface) {
	fmt.Println(121313)
}
