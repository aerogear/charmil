package main

import (
	"github.com/aerogear/charmil/core/commands"
	"log"

	"github.com/aerogear/charmil/examples/plugins/date"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}

	dateCmd, err := date.DateCommand()
	if err != nil {
		log.Fatal(err)
	}

	charmilCommands, err := commands.CharmilCommands()
	if err == nil {
		cmd.AddCommand(charmilCommands)
	}

	cmd.AddCommand(dateCmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
