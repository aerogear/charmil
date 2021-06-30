package example

import (
	"github.com/spf13/cobra"
)

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

	cmd1 := &cobra.Command{
		Use:     "subcmd01",
		Short:   "",
		Example: "examples",
		Run:     func(cmd *cobra.Command, args []string) {},
	}

	cmd2 := &cobra.Command{
		Use:     "subcmd12",
		Short:   "",
		Example: "examples",
		Run:     func(cmd *cobra.Command, args []string) {},
	}

	cmd3 := &cobra.Command{
		Use:     "subcmd03",
		Short:   "",
		Example: "examples mine",
		Run:     func(cmd *cobra.Command, args []string) {},
	}

	cmd1.AddCommand(cmd2)
	cmd.AddCommand(cmd1)
	cmd.AddCommand(cmd3)

	return cmd
}
