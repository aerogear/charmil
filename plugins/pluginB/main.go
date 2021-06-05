package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// RootCmd represents the root command
var RootCmd = &cobra.Command{
	Use:   "pluginB",
	Short: "Plugin B",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Root command of Plugin B called")
	},
}

func main() {
	RootCmd.Execute()
}
