package commands

import (
	"github.com/spf13/cobra"
)

func ValidateCommand() (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:     "validate",
		Short:   "validate custom commands",
		Long:    "validates the custom commands with the defined rules",
		Example: "host charmil validate",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			println("validate")
			return nil;
		},
	}

	return cmd, nil
}
