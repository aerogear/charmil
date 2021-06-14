package logging

import (
	"fmt"
	"io"
	"os"
)

type StdLoggerBuilder struct {
	outStream io.Writer
	errStream io.Writer
}

type StdLogger struct {
	outStream io.Writer
	errStream io.Writer
}

// returns new StdLoggerBuilder object
func NewStdLoggerBuilder() *StdLoggerBuilder {
	builder := new(StdLoggerBuilder)
	return builder
}

// Set input output streams
func (b *StdLoggerBuilder) Streams(out io.Writer, err io.Writer) *StdLoggerBuilder {
	b.outStream = out
	b.errStream = err
	return b
}

// builds a new logger with configuration stored in builder
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

func (l *StdLogger) Output(args ...interface{}) {
	fmt.Fprintln(l.outStream, args...)
}

func (l *StdLogger) Error(args ...interface{}) {
	fmt.Fprintln(l.errStream, args...)
}
