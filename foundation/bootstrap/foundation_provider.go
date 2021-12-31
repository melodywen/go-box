package bootstrap

import (
	"fmt"
	"github.com/melodywen/go-box/contracts/foundation"
)

type FoundationProvider struct {
	app foundation.ApplicationInterface
}

func NewFoundationProvider() *FoundationProvider {
	return &FoundationProvider{}
}

func (foundation *FoundationProvider) Bootstrap(app foundation.ApplicationInterface) {

	fmt.Println(3333)
}
