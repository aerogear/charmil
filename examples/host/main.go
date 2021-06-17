package main

import (
	"log"

	"github.com/aerogear/charmil/pkg/core"
	"github.com/spf13/cobra"

	_ "github.com/aerogear/charmil/examples/plugins/pluginA"
	_ "github.com/namit-chandwani/calculator-charmil-plugin"
)

func main() {
	cmd := &cobra.Command{
		Use:          "hostname",
		Short:        "Host CLI created using Charmil SDK",
		SilenceUsage: true,
	}

	// Add `pluginA` plugin's commands to root
	pluginACmd, err := core.GetRootCmd("pluginA")
	if err != nil {
		log.Fatal(err)
	}
	cmd.AddCommand(pluginACmd)

	// Add `calculator` plugin's commands to root
	calculatorCmd, err := core.GetRootCmd("calculator")
	if err != nil {
		log.Fatal(err)
	}
	cmd.AddCommand(calculatorCmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
