package providers

import (
	"fmt"
	"github.com/melodywen/go-box/illuminate/support"
)

// SchoolServiceProvider providers
type SchoolServiceProvider struct {
	support.ServiceProvider
}

func NewSchoolServiceProvider() *SchoolServiceProvider {
	return &SchoolServiceProvider{}
}

// Boot Bootstrap any application services.
func (provider *SchoolServiceProvider) Boot() {

	provider.App.Make("student")
	fmt.Println("注册 school boot")
}

// Register any application services.
func (provider *SchoolServiceProvider) Register() {
	fmt.Println("注册 school register")
	provider.App.Instance("school", "i am school provider")
	provider.App.Alias("teacher", Teacher{})
}
