package main

import (
	"log"

	"github.com/aerogear/charmil/plugins/calculator"
	"github.com/aerogear/charmil/plugins/hi"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}

	cmd.AddCommand(calculator.AddCommand)
	cmd.AddCommand(hi.HiCommand)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
