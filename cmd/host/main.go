package main

import (
	"log"

	pkg "github.com/aerogear/charmil/pkg/pluginloader"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}

	cmd.AddCommand(pkg.InstallCmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
