package main

import (
	"log"

	"github.com/aerogear/charmil/core/commands"
	echo "github.com/aerogear/charmil/examples/plugin"
	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}

	// Add plugin CLI into host
	echoCmd, err := echo.EchoCommand()
	if err != nil {
		log.Fatal(err)
	}
	root.AddCommand(echoCmd)

	// Add Charmil commands into host
	err = commands.AttachCharmilCommands(root)
	if err != nil {
		log.Fatal(err)
	}

	// Execute root command
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
