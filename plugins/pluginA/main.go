package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd represents the root command
var RootCmd = &cobra.Command{
	Use:   "pluginA",
	Short: "Plugin A",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root command of Plugin A called")
	},
}

func main() {
	RootCmd.Execute()
}
