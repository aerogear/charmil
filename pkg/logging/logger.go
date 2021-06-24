package logging

// Logger is an Interface for logging messages in unified way
type Logger interface {
	// Info prints an formatted output message
	// using fmt.Fprint function
	Info(args ...interface{})
	// Error prints an formatted error message
	// using fmt.Fprint function
	Error(args ...interface{})

	// Infof prints an formatted output message
	// using fmt.Fprintf function
	Infof(format string, args ...interface{})
	// Errorf prints an formatted error message
	// using fmt.Fprintf function
	Errorf(format string, args ...interface{})

	// Infoln prints an formatted output message
	// using fmt.Fprintln function
	Infoln(args ...interface{})

	// Errorln prints an formatted error message
	// using fmt.Fprintln function
	Errorln(args ...interface{})
}
