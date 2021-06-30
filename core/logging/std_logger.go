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
	outStream io.Writer
	errStream io.Writer
}

// StdLogger is a type of logger which uses
// standard output and error streams, or custom writers
type StdLogger struct {
	outStream io.Writer
	errStream io.Writer
}

// NewStdLoggerBuilder returns new StdLoggerBuilder object
func NewStdLoggerBuilder() *StdLoggerBuilder {
	builder := new(StdLoggerBuilder)
	return builder
}

// Streams set input output streams
func (b *StdLoggerBuilder) Streams(out io.Writer, err io.Writer) *StdLoggerBuilder {
	b.outStream = out
	b.errStream = err
	return b
}

// Build creates a new logger instance
// with configuration stored in builder
func (b *StdLoggerBuilder) Build() (logger *StdLogger, err error) {
	logger = new(StdLogger)
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
