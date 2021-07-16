package main

import (
	"log"

	"github.com/aerogear/charmil/core/commands"
	c "github.com/aerogear/charmil/core/config"
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/examples/plugins/adder"
	"github.com/aerogear/charmil/examples/plugins/echo"
	"github.com/spf13/cobra"
)

// Defines the configuration keys of the host CLI.
//
// CONSTRAINT: All fields of the config struct need to be exportable
type config struct {
	Key1    string
	Key2    string
	Key3    string
	Plugins map[string]interface{}
}

// Stores the path of the local config file
const cfgFilePath = "./examples/host/config.json"

var (
	// Stores an instance of the charmil factory
	cmdFactory *factory.Factory

	// Stores an instance of the charmil config handler
	h *c.CfgHandler

	// Initializes a zero-valued struct and stores its address
	cfg = &config{}

	// Root command of the host CLI
	root = &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}
)

func init() {
	// Links the specified local config file path and current config
	// struct pointer to a new config handler instance and returns it
	h = c.NewHandler(cfgFilePath, cfg)

	// Loads config values from the local config file
	err := h.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Creates a new factory instance with default settings
	cmdFactory = factory.Default(nil, h)
}

func main() {

	// Add plugin CLI into host
	echoCmd, err := echo.EchoCommand(cmdFactory)
	if err != nil {
		log.Fatal(err)
	}

	// Stores the root command of the `adder` plugin
	adderCmd, err := adder.AdderCommand(cmdFactory)
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

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
