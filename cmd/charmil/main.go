package main

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "charmil",
		Short:        "Commands for managing plugins",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to Charmil")
		},
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
