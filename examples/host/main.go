package main

import (
	"log"

	"github.com/aerogear/charmil/examples/plugins/calculator"
	"github.com/aerogear/charmil/examples/plugins/date"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}

	cmd.AddCommand(calculator.RootCommand())
	cmd.AddCommand(date.DateCommand())

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
