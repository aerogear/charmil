package main

import (
	"log"

	"github.com/aerogear/charmil/examples/plugins/calculator"
	"github.com/aerogear/charmil/examples/plugins/date"
	"github.com/aerogear/charmil/pkg/factory"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}

	// init default factory for the host
	// pass to all the commands
	defaultFactory := factory.Default()

	cmd.AddCommand(calculator.RootCommand(defaultFactory))
	cmd.AddCommand(date.DateCommand(defaultFactory))

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
