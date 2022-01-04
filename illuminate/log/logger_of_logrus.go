package log

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// LoggerOfLogrus struct
type LoggerOfLogrus struct {
	contextLogger *log.Entry
}

// NewLoggerOfLogrus new instance
func NewLoggerOfLogrus() *LoggerOfLogrus {
	log.SetOutput(os.Stdout)
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	contextLogger := log.WithFields(log.Fields{})
	return &LoggerOfLogrus{
		contextLogger: contextLogger,
	}
}

// Trace Something very low level.
func (l *LoggerOfLogrus) Trace(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Trace(message)
}

// Debug Useful debugging information.
func (l *LoggerOfLogrus) Debug(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Debug(message)
}

// Info Something noteworthy happened!
func (l *LoggerOfLogrus) Info(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Info(message)
}

// Warn You should probably take a look at this.
func (l *LoggerOfLogrus) Warn(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Warn(message)
}

// Error Something failed but I'm not quitting.
func (l *LoggerOfLogrus) Error(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Error(message)
}

// Fatal Calls os.Exit(1) after logging
func (l *LoggerOfLogrus) Fatal(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Fatal(message)
}

// Panic Calls panic() after logging
func (l *LoggerOfLogrus) Panic(message string, fields map[string]interface{}) {
	l.contextLogger.WithFields(fields).Panic(message)
}
