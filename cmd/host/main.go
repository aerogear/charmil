package main

import (
	"log"

	"github.com/aerogear/charmil/pkg/pluginloader"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}

	pluginloader.LoadCommands(cmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
