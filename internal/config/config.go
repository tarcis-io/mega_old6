// Package config loads and provides the application configuration.
package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

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
	EnvLogOutput = "LOG_OUTPUT"

	// EnvServerAddress specifies the environment variable name for configuring the
	// server's address.
	//
	// Expected format: "<host>:port" (e.g., "localhost:8080", ":3000")
	//
	// Default: [DefaultServerAddress]
	EnvServerAddress = "SERVER_ADDRESS"

	// EnvServerReadTimeout specifies the environment variable name for configuring the
	// server's read timeout.
	//
	// Expected format: [time.Duration] (e.g., "5s", "1m")
	//
	// Default: [DefaultServerReadTimeout]
	EnvServerReadTimeout = "SERVER_READ_TIMEOUT"

	// EnvServerReadHeaderTimeout specifies the environment variable name for
	// configuring the server's read header timeout.
	//
	// Expected format: [time.Duration] (e.g., "5s", "1m")
	//
	// Default: [DefaultServerReadHeaderTimeout]
	EnvServerReadHeaderTimeout = "SERVER_READ_HEADER_TIMEOUT"

	// EnvServerWriteTimeout specifies the environment variable name for configuring
	// the server's write timeout.
	//
	// Expected format: [time.Duration] (e.g., "5s", "1m")
	//
	// Default: [DefaultServerWriteTimeout]
	EnvServerWriteTimeout = "SERVER_WRITE_TIMEOUT"

	// EnvServerIdleTimeout specifies the environment variable name for configuring the
	// server's idle timeout.
	//
	// Expected format: [time.Duration] (e.g., "5s", "1m")
	//
	// Default: [DefaultServerIdleTimeout]
	EnvServerIdleTimeout = "SERVER_IDLE_TIMEOUT"

	// EnvServerShutdownTimeout specifies the environment variable name for configuring
	// the server's shutdown timeout.
	//
	// Expected format: [time.Duration] (e.g., "5s", "1m")
	//
	// Default: [DefaultServerShutdownTimeout]
	EnvServerShutdownTimeout = "SERVER_SHUTDOWN_TIMEOUT"
)

const (
	// DefaultLogLevel defines the default [LogLevel], used as the fallback when
	// [EnvLogLevel] is unset.
	DefaultLogLevel LogLevel = LogLevelInfo

	// DefaultLogFormat defines the default [LogFormat], used as the fallback when
	// [EnvLogFormat] is unset.
	DefaultLogFormat LogFormat = LogFormatText

	// DefaultLogOutput defines the default [LogOutput], used as the fallback when
	// [EnvLogOutput] is unset.
	DefaultLogOutput LogOutput = LogOutputStdout

	// DefaultServerAddress defines the default server address, used as the fallback
	// when [EnvServerAddress] is unset.
	DefaultServerAddress = "localhost:8080"

	// DefaultServerReadTimeout defines the default server read timeout, used as the
	// fallback when [EnvServerReadTimeout] is unset.
	DefaultServerReadTimeout = 5 * time.Second

	// DefaultServerReadHeaderTimeout defines the default server read header timeout,
	// used as the fallback when [EnvServerReadHeaderTimeout] is unset.
	DefaultServerReadHeaderTimeout = 2 * time.Second

	// DefaultServerWriteTimeout defines the default server write timeout, used as the
	// fallback when [EnvServerWriteTimeout] is unset.
	DefaultServerWriteTimeout = 10 * time.Second

	// DefaultServerIdleTimeout defines the default server idle timeout, used as the
	// fallback when [EnvServerIdleTimeout] is unset.
	DefaultServerIdleTimeout = 1 * time.Minute

	// DefaultServerShutdownTimeout defines the default server shutdown timeout, used
	// as the fallback when [EnvServerShutdownTimeout] is unset.
	DefaultServerShutdownTimeout = 15 * time.Second
)

const (
	// TCPPortMin defines the minimum port number for TCP connections.
	TCPPortMin = 0

	// TCPPortMax defines the maximum port number for TCP connections.
	TCPPortMax = 65535
)

type (
	Config struct {
		logLevel                LogLevel
		logFormat               LogFormat
		logOutput               LogOutput
		serverAddress           string
		serverReadTimeout       time.Duration
		serverReadHeaderTimeout time.Duration
		serverWriteTimeout      time.Duration
		serverIdleTimeout       time.Duration
		serverShutdownTimeout   time.Duration
	}
)

func New() (*Config, error) {
	l := newLoader()
	cfg := &Config{
		logLevel:  l.logLevel(),
		logFormat: l.logFormat(),
		logOutput: l.logOutput(),
	}
	if err := l.Err(); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return cfg, nil
}

type (
	loader struct {
		errs []error
	}
)

func newLoader() *loader {
	return &loader{}
}

func (l *loader) logLevel() LogLevel {
	env := getEnv(EnvLogLevel, string(DefaultLogLevel))
	switch val := LogLevel(strings.ToLower(strings.TrimSpace(env))); val {
	case LogLevelDebug, LogLevelInfo, LogLevelWarn, LogLevelError:
		return val
	}
	l.appendError(fmt.Errorf("invalid log level (%s) got=%q", EnvLogLevel, env))
	return ""
}

func (l *loader) logFormat() LogFormat {
	env := getEnv(EnvLogFormat, string(DefaultLogFormat))
	switch val := LogFormat(strings.ToLower(strings.TrimSpace(env))); val {
	case LogFormatText, LogFormatJSON:
		return val
	}
	l.appendError(fmt.Errorf("invalid log format (%s) got=%q", EnvLogFormat, env))
	return ""
}

func (l *loader) logOutput() LogOutput {
	env := getEnv(EnvLogOutput, string(DefaultLogOutput))
	val := strings.TrimSpace(env)
	switch v := LogOutput(strings.ToLower(val)); v {
	case LogOutputStdout, LogOutputStderr:
		return v
	}
	if val == "" {
		l.appendError(fmt.Errorf("invalid log output (%s) got=%q", EnvLogOutput, env))
		return ""
	}
	return LogOutput(val)
}

func (l *loader) appendError(err error) {
	l.errs = append(l.errs, err)
}

func (l *loader) Err() error {
	if len(l.errs) == 0 {
		return nil
	}
	return errors.Join(l.errs...)
}

func getEnv(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defaultValue
}
