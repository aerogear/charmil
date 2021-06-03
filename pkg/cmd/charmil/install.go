package charmil

import (
	"fmt"

	"github.com/spf13/cobra"
)

// InstallCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install plugin(s)",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("install called")
		// Temporary code. Will be later replaced with logic which downloads, extracts, installs and validates binary from manifest file

		return nil
	},
}

func init() {
}
