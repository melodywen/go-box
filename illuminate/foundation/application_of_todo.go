package foundation

func (app *Application) Environment() {
	panic("implement me")
}

func (app *Application) RunningInConsole() {
	panic("implement me")
}

func (app *Application) RunningUnitTests() {
	panic("implement me")
}

func (app *Application) IsDownForMaintenance() {
	panic("implement me")
}

// ResolveProvider Resolve a service provider instance from the class name.
func (app *Application) ResolveProvider() {
	panic("implement me")
}

func (app *Application) GetLocale() {
	panic("implement me")
}

func (app *Application) GetNamespace() {
	panic("implement me")
}

func (app *Application) SetLocale() {
	panic("implement me")
}

func (app *Application) ShouldSkipMiddleware() {
	panic("implement me")
}

func (app *Application) Terminate() {
	panic("implement me")
}
