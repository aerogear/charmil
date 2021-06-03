package charmil

import (
	"fmt"

	"github.com/spf13/cobra"
)

// DeactivateCmd represents the deactivate command
var DeactivateCmd = &cobra.Command{
	Use:   "deactivate",
	Short: "Deactivate plugin(s)",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deactivate called")
	},
}

func init() {
}
