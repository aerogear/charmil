package main

import (
	"log"

	"github.com/aerogear/charmil/core/commands"
	"github.com/aerogear/charmil/core/config"
	echo "github.com/aerogear/charmil/examples/plugin"
	"github.com/aerogear/charmil/examples/plugins/adder"
	"github.com/spf13/cobra"
)

var (
	h *config.Handler

	f = config.File{
		Name: "config",
		Type: "yaml",
		Path: "./examples/host",
	}

	root = &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}
)

func init() {
	h = config.New()

	h.InitFile(f)

	err := h.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	h.SetValue("key4", "val4")

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
	adderCmd, adderCfg, err := adder.AdderCommand()
	if err != nil {
		log.Fatal(err)
	}

	root.AddCommand(adderCmd)

	h.SetPluginCfg("adder", adderCfg)

	h.MergePluginCfg()

	err = h.Save()
	if err != nil {
		log.Fatal(err)
	}

	// Execute root command
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
