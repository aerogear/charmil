package validator

import (
	"github.com/spf13/cobra"
)

// ValidationError is a default validation error
type ValidationError struct {
	Name string
	Err  error
	Rule string
	Cmd  *cobra.Command
}

// StatusLog is used for providing info
// about validation of commands
type StatusLog struct {
	// totalTested represents number of commands checked
	TotalTested int
	// totalErrors reperesents number of errors
	TotalErrors int
	// errors in command
	Errors []ValidationError
}
