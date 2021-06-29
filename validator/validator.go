package validator

import (
	"fmt"

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

// Traverse is used to traverse and validate
// the command and it's descendant commands
func Traverse(cmd *cobra.Command, verbose bool, info StatusLog, validate func(cmd *cobra.Command, verbose bool) []ValidationError) []ValidationError {
	// validate the root command
	err := validate(cmd, verbose)
	// record stats
	info.TotalTested++
	info.TotalErrors += len(err)
	info.Errors = append(info.Errors, err...)

	// traverse descendents of cmd
	for _, child := range cmd.Commands() {

		// base case
		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}

		// recursive call for ValidateHelper
		if err := Traverse(child, verbose, info, validate); err != nil {
			return err
		}
	}

	// prints additional info in debug mode
	if verbose {
		fmt.Printf("commands checked: %d\nchecks failed: %d\n", info.TotalTested, info.TotalErrors)
	}

	return info.Errors
}
