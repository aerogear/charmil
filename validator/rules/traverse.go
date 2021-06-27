package rules

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Traverse is used to traverse and validate
// the command and it's descendant commands
func Traverse(cmd *cobra.Command, verbose bool, info StatusLog, x interface{}, validate func(cmd *cobra.Command, verbose bool) []ValidationError) []ValidationError {
	// validate the root command
	err := validate(cmd, verbose)
	// record stats
	info.totalTested++
	info.totalErrors += len(err)
	info.errors = append(info.errors, err...)

	// traverse descendents of cmd
	for _, child := range cmd.Commands() {

		// base case
		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}

		// recursive call for ValidateHelper
		if err := Traverse(child, verbose, info, x, validate); err != nil {
			return err
		}
	}

	// prints additional info in debug mode
	if verbose {
		fmt.Printf("commands checked: %d\nchecks failed: %d\n", info.totalTested, info.totalErrors)
	}

	return info.errors
}
