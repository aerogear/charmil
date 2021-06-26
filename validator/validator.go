package validator

import (
	"github.com/spf13/cobra"
)

type Rule interface {
	Validate(cmd *cobra.Command, verbose bool) []Error
}
