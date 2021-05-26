package extensions

import (
	"fmt"

	"github.com/aerogear/charmil/cmd/extensions/install"
	"github.com/aerogear/charmil/cmd/extensions/installed"
	"github.com/aerogear/charmil/cmd/extensions/list"
	"github.com/aerogear/charmil/cmd/extensions/remove"
	"github.com/spf13/cobra"
)

// ExtensionsCmd represents the extensions command
var ExtensionsCmd = &cobra.Command{
	Use:   "extensions",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("extensions called")
	},
}

func init() {
	ExtensionsCmd.AddCommand(install.InstallCmd, installed.InstalledCmd, list.ListCmd, remove.RemoveCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ExtensionsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ExtensionsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
