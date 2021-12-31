package http

type KernelInterface interface {

	// Bootstrap the application for HTTP requests.
	Bootstrap()

	// Handle  init kernel
	Handle()

	// Terminate Perform any final actions for the request lifecycle.
	Terminate()

	// GetApplication Get the Laravel application instance.
	GetApplication()
}
