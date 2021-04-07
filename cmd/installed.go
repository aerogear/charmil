package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var installedCmd = &cobra.Command{
	Use:   "List Installed Extensions",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("installed called")
	},
}

func init() {
	rootCmd.AddCommand(installedCmd)
}
