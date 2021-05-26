package extensions

import (
	"fmt"

	"github.com/aerogear/charmil/cmd/extensions/install"
	"github.com/aerogear/charmil/cmd/extensions/installed"
	"github.com/aerogear/charmil/cmd/extensions/list"
	"github.com/aerogear/charmil/cmd/extensions/remove"
	"github.com/spf13/cobra"
)

var ExtensionsCmd = &cobra.Command{
	Use:   "extensions",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("extensions called")
	},
}

func init() {
	ExtensionsCmd.AddCommand(install.InstallCmd, installed.InstalledCmd, list.ListCmd, remove.RemoveCmd)
}
