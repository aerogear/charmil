package validator

import (
	"github.com/spf13/cobra"
)

// RuleOptions is present in each rule
// to control the options limited to the rule
type RuleOptions struct {
	Disable      bool
	Verbose      bool
	SkipCommands map[string]bool
}

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
