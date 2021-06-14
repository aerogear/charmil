package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "hostname",
		Short:        "Host CLI made using Charmil SDK",
		SilenceUsage: true,
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
