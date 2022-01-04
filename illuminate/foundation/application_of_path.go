package foundation

import (
	"path"
)

// setBasePath Set the base path for the application.
func (app *Application) setBasePath(basePath string) {
	app.basePath = basePath
	app.bindPathsInContainer()
}

//bindPathsInContainer Bind all the application paths in the container.
func (app *Application) bindPathsInContainer() {
	app.Instance("path", app.path(""))
	app.Instance("path.base", app.BasePath(""))
	app.Instance("path.config", app.ConfigPath(""))
	app.Instance("path.storage", app.StoragePath())
	app.Instance("path.database", app.DatabasePath(""))
	app.Instance("path.resources", app.ResourcePath(""))
	app.Instance("path.bootstrap", app.BootstrapPath(""))
}

//path Get the path to the application "app" directory.
func (app *Application) path(pathSuffix string) string {
	appPath := path.Join(app.basePath, pathSuffix)
	return appPath
}

// BasePath Get the base path of the Laravel installation.
func (app *Application) BasePath(pathSuffix string) string {
	appPath := path.Join(app.basePath, pathSuffix)
	return appPath
}

// BootstrapPath Get the path to the bootstrap directory.
func (app *Application) BootstrapPath(pathSuffix string) string {
	return path.Join(app.basePath, "bootstrap", pathSuffix)
}

// ConfigPath Get the path to the application configuration files.
func (app *Application) ConfigPath(pathSuffix string) string {
	return path.Join(app.basePath, "config", pathSuffix)
}

// DatabasePath Get the path to the database directory.
func (app *Application) DatabasePath(pathSuffix string) string {
	return path.Join(app.basePath, "database", pathSuffix)
}

// ResourcePath Get the path to the resources directory.
func (app *Application) ResourcePath(pathSuffix string) string {
	return path.Join(app.basePath, "resources", pathSuffix)
}

// StoragePath Get the path to the storage directory.
func (app *Application) StoragePath() string {
	return path.Join(app.basePath, "storage")
}
