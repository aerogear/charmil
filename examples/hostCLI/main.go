package main

import (
	"log"

	"github.com/aerogear/charmil/examples/plugins/pluginA"
	"github.com/aerogear/charmil/pkg/core"
	"github.com/aerogear/charmil/pkg/factory"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "hostname",
		Short:        "Host CLI made using Charmil SDK",
		SilenceUsage: true,
	}
	defaultFactory := factory.Default()
	p := pluginA.GetPlugin()

	cmd.AddCommand(core.GetPluginRootCmd(p, defaultFactory))

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
