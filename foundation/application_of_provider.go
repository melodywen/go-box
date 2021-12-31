package foundation


// Boot the application's service providers.
func (app *Application) Boot() {
	if app.isBooted() {
		return
	}

	// Once the application has booted we will also fire some "booted" callbacks
	// for any listeners that need to do work after this initial booting gets
	// finished. This is useful when ordering the boot-up processes we run.
}
//  Determine if the application has booted.
func (app *Application) isBooted()  bool{
	return app.booted
}
