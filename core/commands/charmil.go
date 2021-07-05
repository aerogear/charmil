package commands

import (
	"github.com/spf13/cobra"
)

func AttachCharmilCommands(hostRoot *cobra.Command) error {

	if !inDev {
		return nil
	}

	// Placeholder command. To be replaced later by actual commands
	cmd := &cobra.Command{
		Use:          "charmil",
		Short:        "built in charmil commands",
		SilenceUsage: false,
	}

	hostRoot.AddCommand(cmd)
	return nil
}
