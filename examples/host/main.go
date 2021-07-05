package main

import (
	"github.com/aerogear/charmil/core/commands"
	"github.com/aerogear/charmil/examples/plugins/date"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	root := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}

	dateCmd, err := date.DateCommand()
	if err != nil {
		log.Fatal(err)
	}

	root.AddCommand(dateCmd)

	err = commands.AttachCharmilCommands(root)
	if err != nil {
		log.Fatal(err)
	}

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
