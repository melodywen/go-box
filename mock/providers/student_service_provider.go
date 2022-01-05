package providers

import (
	"fmt"
	"github.com/melodywen/go-box/illuminate/support"
)

// StudentServiceProvider providers
type StudentServiceProvider struct {
	support.ServiceProvider
}

func NewStudentServiceProvider() *StudentServiceProvider {
	return &StudentServiceProvider{}
}

// Boot Bootstrap any application services.
func (provider *StudentServiceProvider) Boot() {
	fmt.Println("注册 student boot")
}

// Register any application services.
func (provider *StudentServiceProvider) Register() {
	fmt.Println("注册student register")

	provider.App.Bind("student", func(teacher Teacher) {
		fmt.Println(teacher)
	}, true)


}
