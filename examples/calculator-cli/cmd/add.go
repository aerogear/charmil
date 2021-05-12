package cmd

import "charmil/internal/add"

func init() {
	rootCmd.AddCommand(add.AddCmd)
}
