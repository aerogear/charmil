package main

import (
	"log"

	"github.com/aerogear/charmil/pkg/cmd/charmil"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "charmil",
		Short:        "Commands for managing plugins",
		SilenceUsage: true,
	}

	cmd.AddCommand(charmil.InstallCmd)

	// err = doc.GenMarkdownTree(cmd, "./docs/commands")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
