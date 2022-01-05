package providers

import (
	"fmt"
	"github.com/melodywen/go-box/illuminate/support"
)

// TeacherServiceProvider providers
type TeacherServiceProvider struct {
	support.ServiceProvider
}

// Teacher a teacher
type Teacher struct {
	name     string
	describe string
}

// NewTeacherServiceProvider new instance
func NewTeacherServiceProvider() *TeacherServiceProvider {
	return &TeacherServiceProvider{}
}

// Boot Bootstrap any application services.
func (provider *TeacherServiceProvider) Boot() {
	fmt.Println("注册 teacher boot")
}

// Register any application services.
func (provider *TeacherServiceProvider) Register() {
	fmt.Println("注册 teacher register")

	provider.App.Instance("teacher", Teacher{
		name:     "tom",
		describe: "i am teacher",
	})
}
