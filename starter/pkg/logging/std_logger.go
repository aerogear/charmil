// This file contains a logger that uses the standard output and error streams, or custom writers.

package logging

import (
	"fmt"
	"io"
	"os"
)

// StdLoggerBuilder contains the configuration and logic needed to build a logger that uses the
// standard output and error streams, or custom writers.
type StdLoggerBuilder struct {
	debugEnabled bool
	infoEnabled  bool
	errorEnabled bool
	outStream    io.Writer
	errStream    io.Writer
}

// StdLogger is a logger that uses the standard output and error streams, or custom writers.
type StdLogger struct {
	debugEnabled bool
	infoEnabled  bool
	errorEnabled bool
	outStream    io.Writer
	errStream    io.Writer
}

// NewStdLoggerBuilder creates a builder that knows how to build a logger that uses the standard
// output and error streams, or custom writers. By default these loggers will have enabled the
// information, warning and error levels
func NewStdLoggerBuilder() *StdLoggerBuilder {
	// Allocate the object:
	builder := new(StdLoggerBuilder)

	// Set default values:
	builder.debugEnabled = false
	builder.infoEnabled = true
	builder.errorEnabled = true

	return builder
}

// Streams sets the standard output and error streams to use. If not used then the logger will use
// os.Stdout and os.Stderr.
func (b *StdLoggerBuilder) Streams(out io.Writer, err io.Writer) *StdLoggerBuilder {
	b.outStream = out
	b.errStream = err
	return b
}

// Debug enables or disables the debug level.
func (b *StdLoggerBuilder) Debug(flag bool) *StdLoggerBuilder {
	b.debugEnabled = flag
	return b
}

// Info enables or disables the information level.
func (b *StdLoggerBuilder) Info(flag bool) *StdLoggerBuilder {
	b.infoEnabled = flag
	return b
}

// Error enables or disables the error level.
func (b *StdLoggerBuilder) Error(flag bool) *StdLoggerBuilder {
	b.errorEnabled = flag
	return b
}

// Build creates a new logger using the configuration stored in the builder.
func (b *StdLoggerBuilder) Build() (logger *StdLogger, err error) {
	// Allocate and populate the object:
	logger = new(StdLogger)
	logger.debugEnabled = b.debugEnabled
	logger.infoEnabled = b.infoEnabled
	logger.errorEnabled = b.errorEnabled
	logger.outStream = b.outStream
	logger.errStream = b.errStream
	if logger.outStream == nil {
		logger.outStream = os.Stdout
	}
	if logger.errStream == nil {
		logger.errStream = os.Stderr
	}

	return
}

// DebugEnabled returns true iff the debug level is enabled.
func (l *StdLogger) DebugEnabled() bool {
	return l.debugEnabled
}

// InfoEnabled returns true iff the information level is enabled.
func (l *StdLogger) InfoEnabled() bool {
	return l.infoEnabled
}

// ErrorEnabled returns true iff the error level is enabled.
func (l *StdLogger) ErrorEnabled() bool {
	return l.errorEnabled
}

// Info sends to the log an information message formatted using the fmt.Fprintln function and the
// given arguments.
func (l *StdLogger) Info(args ...interface{}) {
	if l.infoEnabled {
		fmt.Fprintln(l.errStream, args...)
	}
}

// Infof sends to the log an information message formatted using the fmt.Fprintf function and the
// given format and arguments.
func (l *StdLogger) Infof(format string, args ...interface{}) {
	if l.infoEnabled {
		fmt.Fprintf(l.errStream, format+"\n", args...)
	}
}

// Debug sends to the log a debug message formatted using the fmt.Fprintln function and the given
// arguments.
func (l *StdLogger) Debug(args ...interface{}) {
	if l.debugEnabled {
		fmt.Fprintln(l.errStream, args...)
	}
}

// Debugf sends to the log a debug message formatted using the fmt.Fprintf function and the given
// format and arguments.
func (l *StdLogger) Debugf(format string, args ...interface{}) {
	if l.debugEnabled {
		fmt.Fprintf(l.errStream, format+"\n", args...)
	}
}

// Error sends to the log an error message formatted using the fmt.Fprintln function and the given
// arguments.
func (l *StdLogger) Error(args ...interface{}) {
	if l.infoEnabled {
		fmt.Fprintln(l.errStream, args...)
	}
}

// Errorf sends to the log an error message formatted using the fmt.Fprintf function and the given
// format and arguments.
func (l *StdLogger) Errorf(format string, args ...interface{}) {
	if l.infoEnabled {
		fmt.Fprintf(l.errStream, format+"\n", args...)
	}
}
