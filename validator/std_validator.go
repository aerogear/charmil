package validator

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Validate validates the cobra commands
// and all the descendants
func Validate(cmd *cobra.Command) error {
	err := handleDefaultErrors(cmd)
	if err != nil {
		return err
	}

	// handeling errors for subcommands of a command
	subCommmands := cmd.Commands()
	for _, child := range subCommmands {
		err := handleDefaultErrors(child)
		if err != nil {
			return err
		}
	}

	return nil
}

// ValidateCustom will allow user to
// write their own custom rules for validation
// over the default rules
func ValidateCustom(cmd *cobra.Command, handleErrors func() error) error {
	err1 := Validate(cmd)
	if err1 != nil {
		return err1
	}
	err2 := handleErrors()
	if err2 != nil {
		return err2
	}

	return nil
}

// handleDefaultErrors handles the default errors
func handleDefaultErrors(cmd *cobra.Command) error {
	cmdPath := cmd.CommandPath()

	if len(cmd.Use) < 1 {
		return fmt.Errorf("%s: length of cmd.Use should be more than 1", cmdPath)
	}

	if len(cmd.Short) < 5 {
		return fmt.Errorf("%s: length of cmd.Short should be more than 5", cmdPath)
	}

	if len(cmd.Long) < 10 {
		return fmt.Errorf("%s: length of cmd.Long should be more than 10", cmdPath)
	}

	if len(cmd.Example) < 1 {
		return fmt.Errorf("%s: length of cmd.Example should be more than 1", cmdPath)
	}

	if cmd.Args == nil {
		return fmt.Errorf("%s: provide Args function to the command", cmdPath)
	}

	return nil
}
