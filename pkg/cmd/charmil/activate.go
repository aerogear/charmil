package charmil

import (
	"fmt"

	"github.com/spf13/cobra"
)

// ActivateCmd represents the activate command
var ActivateCmd = &cobra.Command{
	Use:   "activate",
	Short: "Activate plugin(s)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("activate called")

		// Todo: Add code for plugin validation
	},
}

func init() {
}
