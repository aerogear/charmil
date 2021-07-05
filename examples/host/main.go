package main

import (
	"log"

	"github.com/aerogear/charmil/core/commands"
	c "github.com/aerogear/charmil/core/config"
	echo "github.com/aerogear/charmil/examples/plugin"
	"github.com/aerogear/charmil/examples/plugins/adder"
	"github.com/spf13/cobra"
)

// CONSTRAINT: All fields of config struct need to be exportable
type config struct {
	Key1 string
	Key2 string
	Key3 string
	Key4 string
}

const cfgFilePath = "./examples/host/config.json"

var (
	// Stores an instance of the charmil config handler
	h *c.Handler

	cfg = &config{}

	// Root command of the host CLI
	root = &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}
)

func init() {
	// Assigns a new instance of the charmil config handler
	// Links the handler instance to a local config file
	h = c.NewHandler(cfgFilePath, cfg)

	// Loads config values from the local config file
	err := h.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// TODO: Initialize the factory struct here and pass it to the plugins

	// Sets a value into config
	cfg.Key4 = "val4"

	// Add plugin CLI into host
	echoCmd, err := echo.EchoCommand()
	if err != nil {
		log.Fatal(err)
	}

	// Stores the root command and the config map of the `adder` plugin
	adderCmd, err := adder.AdderCommand(cfgFilePath)
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

	// Writes the current config into the local config file
	err = h.Save()
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println("Viper host settings are: ", h.GetAllSettings())

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
