// Package config loads and provides the application configuration.
package config

type (
	// LogLevel represents the severity or verbosity of log records.
	LogLevel string
)

const (
	// LogLevelDebug captures detailed information, typically useful for development
	// and debugging.
	LogLevelDebug LogLevel = "debug"

	// LogLevelInfo captures general information about the application's operation.
	LogLevelInfo LogLevel = "info"

	// LogLevelWarn captures non-critical events or potentially harmful situations.
	LogLevelWarn LogLevel = "warn"

	// LogLevelError captures critical events or errors that require immediate
	// attention.
	LogLevelError LogLevel = "error"
)

type (
	// LogFormat represents the encoding style of log records.
	LogFormat string
)

const (
	LogFormatText LogFormat = "text"
	LogFormatJSON LogFormat = "json"
)

type (
	// LogOutput represents the destination stream of log records.
	LogOutput string
)

const (
	LogOutputStdout LogOutput = "stdout"
	LogOutputStderr LogOutput = "stderr"
)
