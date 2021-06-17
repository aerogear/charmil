package main

import (
	"log"

	_ "github.com/aerogear/charmil/examples/plugins/pluginA"
	"github.com/aerogear/charmil/pkg/core"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "hostname",
		Short:        "Host CLI created using Charmil SDK",
		SilenceUsage: true,
	}

	pluginCmd, err := core.GetRootCmd("pluginA")
	if err != nil {
		log.Fatal(err)
	}
	cmd.AddCommand(pluginCmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
