package logging

type Logger interface {
	// DebugEnabled returns true if the debug level is enabled.
	DebugEnabled() bool

	// InfoEnabled returns true if the information level is enabled.
	InfoEnabled() bool

	// Debug sends to the log a debug message formatted using the fmt.Sprintf function and the
	// given format and arguments.
	Debug(args ...interface{})

	// Debugf sends to the log a debug message formatted using the fmt.Sprintf function and the
	// given format and arguments.
	Debugf(format string, args ...interface{})

	// Info sends to the log an information message formatted using the fmt.Sprintf function and
	// the given format and arguments.
	Info(args ...interface{})

	// Infof sends to the log an information message formatted using the fmt.Sprintf function and
	// the given format and arguments.
	Infof(format string, args ...interface{})

	// Error sends to the log an error message formatted using the fmt.Sprintln function and the
	// given format and arguments.
	Error(args ...interface{})

	// Errorf sends to the log an error message formatted using the fmt.Sprint function and the
	// given format and arguments.
	Errorf(format string, args ...interface{})
}
