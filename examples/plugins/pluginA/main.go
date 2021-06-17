package pluginA

import (
	"fmt"

	"github.com/aerogear/charmil/pkg/core"
	"github.com/spf13/cobra"
)

type myPlugin struct{}

// This will be executed when the host imports this package in an unnamed manner
func init() {
	core.Register("pluginA", &myPlugin{})
}

func (p myPlugin) CreateRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "pluginA",
		Short: "Plugin A",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Root command of Plugin A called")
		},
	}
	return rootCmd
}
