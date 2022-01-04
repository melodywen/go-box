package log

// LoggerInterface logger abstract
type LoggerInterface interface {
	// Trace Something very low level.
	Trace(message string, fields map[string]interface{})
	// Debug Useful debugging information.
	Debug(message string, fields map[string]interface{})
	// Info Something noteworthy happened!
	Info(message string, fields map[string]interface{})
	// Warn You should probably take a look at this.
	Warn(message string, fields map[string]interface{})
	// Error Something failed but I'm not quitting.
	Error(message string, fields map[string]interface{})
	// Fatal Calls os.Exit(1) after logging
	Fatal(message string, fields map[string]interface{})
	// Panic Calls panic() after logging
	Panic(message string, fields map[string]interface{})
}
