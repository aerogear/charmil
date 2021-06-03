package charmil

import (
	"fmt"

	"github.com/spf13/cobra"
)

// InstalledCmd represents the installed command
var InstalledCmd = &cobra.Command{
	Use:   "installed",
	Short: "Show installed plugins",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("installed called")
	},
}

func init() {
}
