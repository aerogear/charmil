package logging

import (
	"fmt"
	"io"
	"os"
)

// Standard Logger Implementation

// StdLoggerBuilder is a type which contains
// the configuration to build the logger which uses
// standard output and error streams or custom writers
type StdLoggerBuilder struct {
	debugEnabled bool
	infoEnabled  bool
	errorEnabled bool
	outStream    io.Writer
	errStream    io.Writer
}

// StdLogger is a type of logger which uses
// standard output and error streams, or custom writers
type StdLogger struct {
	debugEnabled bool
	infoEnabled  bool
	errorEnabled bool
	outStream    io.Writer
	errStream    io.Writer
}

// NewStdLoggerBuilder returns new StdLoggerBuilder object
func NewStdLoggerBuilder() *StdLoggerBuilder {
	// Allocate the object:
	builder := new(StdLoggerBuilder)

	// Set default values:
	builder.debugEnabled = false
	builder.infoEnabled = true
	builder.errorEnabled = true

	return builder
}

// Streams set input output streams
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

// Build creates a new logger instance
// with configuration stored in builder
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

// Info prints an formatted output message
// using fmt.Fprint function
func (l *StdLogger) Info(args ...interface{}) {
	fmt.Fprint(l.outStream, args...)
}

// Infof prints an formatted output message
// using fmt.Fprintf function
func (l *StdLogger) Infof(format string, args ...interface{}) {
	fmt.Fprintf(l.outStream, format, args...)
}

// Infoln prints an formatted output message
// using fmt.Fprintln function
func (l *StdLogger) Infoln(args ...interface{}) {
	fmt.Fprintln(l.outStream, args...)
}

// Error prints an formatted error message
// using fmt.Fprint function
func (l *StdLogger) Error(args ...interface{}) {
	fmt.Fprint(l.errStream, args...)
}

// Errorf prints an formatted error message
// using fmt.Fprintf function
func (l *StdLogger) Errorf(format string, args ...interface{}) {
	fmt.Fprintf(l.errStream, format, args...)
}

// Errorln prints an formatted error message
// using fmt.Fprintln function
func (l *StdLogger) Errorln(args ...interface{}) {
	fmt.Fprintln(l.errStream, args...)
}
