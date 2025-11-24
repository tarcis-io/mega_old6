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
	// LogFormatText renders log records as human-readable text.
	LogFormatText LogFormat = "text"

	// LogFormatJSON renders log records as structured JSON objects.
	LogFormatJSON LogFormat = "json"
)

type (
	// LogOutput represents the destination stream of log records.
	LogOutput string
)

const (
	// LogOutputStdout writes log records to the standard output stream (stdout).
	LogOutputStdout LogOutput = "stdout"

	// LogOutputStderr writes log records to the standard error stream (stderr).
	LogOutputStderr LogOutput = "stderr"
)

const (
	// EnvLogLevel specifies the environment variable name for configuring the
	// [LogLevel].
	//
	// Expected values:
	//
	//  - [LogLevelDebug]
	//  - [LogLevelInfo]
	//  - [LogLevelWarn]
	//  - [LogLevelError]
	//
	// Default: [DefaultLogLevel]
	EnvLogLevel = "LOG_LEVEL"

	// EnvLogFormat specifies the environment variable name for configuring the
	// [LogFormat].
	//
	// Expected values:
	//
	//  - [LogFormatText]
	//  - [LogFormatJSON]
	//
	// Default: [DefaultLogFormat]
	EnvLogFormat = "LOG_FORMAT"

	// EnvLogOutput specifies the environment variable name for configuring the
	// [LogOutput].
	//
	// Expected values:
	//
	//  - [LogOutputStdout]
	//  - [LogOutputStderr]
	//  - A custom string (typically a file path)
	//
	// Default: [DefaultLogOutput]
	EnvLogOutput               = "LOG_OUTPUT"
	EnvServerAddress           = "SERVER_ADDRESS"
	EnvServerReadTimeout       = "SERVER_READ_TIMEOUT"
	EnvServerReadHeaderTimeout = "SERVER_READ_HEADER_TIMEOUT"
	EnvServerWriteTimeout      = "SERVER_WRITE_TIMEOUT"
	EnvServerIdleTimeout       = "SERVER_IDLE_TIMEOUT"
	EnvServerShutdownTimeout   = "SERVER_SHUTDOWN_TIMEOUT"
)
