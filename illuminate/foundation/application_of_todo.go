package foundation

// Environment Get or check the current application environment.
func (app *Application) Environment() {
	panic("implement me")
}

// RunningInConsole Determine if the application is running in the console.
func (app *Application) RunningInConsole() {
	panic("implement me")
}

// RunningUnitTests Determine if the application is running unit tests.
func (app *Application) RunningUnitTests() {
	panic("implement me")
}

// IsDownForMaintenance Determine if the application is currently down for maintenance.
func (app *Application) IsDownForMaintenance() {
	panic("implement me")
}

// ResolveProvider Resolve a service provider instance from the class name.
func (app *Application) ResolveProvider() {
	panic("implement me")
}

// GetLocale Get the current application locale.
func (app *Application) GetLocale() {
	panic("implement me")
}

// GetNamespace Get the application namespace.
func (app *Application) GetNamespace() {
	panic("implement me")
}

// SetLocale Set the current application locale.
func (app *Application) SetLocale() {
	panic("implement me")
}

// ShouldSkipMiddleware Determine if middleware has been disabled for the application.
func (app *Application) ShouldSkipMiddleware() {
	panic("implement me")
}

// Terminate the application.
func (app *Application) Terminate() {
	panic("implement me")
}
