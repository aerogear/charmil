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
	// Stores an instance of the charmil config handler
	h *config.Handler

	// Stores the local config file settings
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
	// Assigns a new instance of the charmil config handler
	h = config.New()

	// Links the handler instance to a local config file
	h.InitFile(f)

	// Loads config values from the local config file
	err := h.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Sets a dummy value into config
	h.SetValue("key4", "val4")

	// Add plugin CLI into host
	echoCmd, err := echo.EchoCommand()
	if err != nil {
		log.Fatal(err)
	}

	// Stores the root command and the config map of the `adder` plugin
	adderCmd, adderCfg, err := adder.AdderCommand()
	if err != nil {
		log.Fatal(err)
	}

	root.AddCommand(echoCmd)
	root.AddCommand(adderCmd)

	// Add Charmil commands into host
	err = commands.AttachCharmilCommands(root)
	if err != nil {
		log.Fatal(err)
	}

	// Maps the plugin name to its imported config map
	h.SetPluginCfg("adder", adderCfg)

	// Stores config of every imported plugin into the current config
	h.MergePluginCfg()

	// Writes the current config into the local config file
	err = h.Save()
	if err != nil {
		log.Fatal(err)
	}

	// Execute root command
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
