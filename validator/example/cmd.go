package example

import (
	"github.com/spf13/cobra"
)

type command struct {
	name    *cobra.Command
	use     string
	short   string
	example string
}

func NewCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "cmd0",
		Short:   "This is the short",
		Long:    "This is long",
		Example: "My Example",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	var commands []command = []command{
		{use: "subcmd01", short: "short01", example: "example01"},
		{use: "subcmd12", short: "short12", example: "example12"},
		{use: "subcmd03", short: "short03", example: "example03"},
	}

	for _, cm := range commands {
		cm.name = &cobra.Command{
			Use:     cm.use,
			Short:   cm.short,
			Example: cm.example,
			Run:     func(cmd *cobra.Command, args []string) {},
		}
		cmd.AddCommand(cm.name)
	}

	return cmd
}
