package validator

import (
	"github.com/aerogear/charmil/validator/rules"
	"github.com/spf13/cobra"
)

// Rule is the interface which is implemented
// by every rule defined in validator package
type Rule interface {
	// Validate validates the cobra command
	// and returns errors according to the provided rules
	// verbose option is present which prints debug info
	// Validate(cmd *cobra.Command, verbose bool) []rules.Error
	ValidateLength(cmd *cobra.Command, verbose bool) []rules.ValidationError
	ValidateMustPresent(cmd *cobra.Command, verbose bool) []rules.ValidationError
}
