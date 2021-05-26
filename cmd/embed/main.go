package main

import (
	"log"

	"github.com/aerogear/charmil/pkg/pluginloader"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "embed",
		Short:        "embeding commands",
		SilenceUsage: true,
	}

	err := pluginloader.AddCommands(cmd)
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
