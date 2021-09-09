package main

import (
	"context"
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/aerogear/charmil/core/config"
	"github.com/aerogear/charmil/core/utils/localize"
	"github.com/aerogear/charmil/starter/internal/update"
	"github.com/aerogear/charmil/starter/pkg/factory"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"golang.org/x/text/language"

	"github.com/aerogear/charmil/starter/pkg/cmd/root"
)

// Defines the configuration keys of the host CLI.
// CONSTRAINT: All fields of the options struct need to be exportable
type options struct {
	LocConfig localize.Config
}

// Stores the path of the local config file
const cfgFilePath = "./config.json"

var (
	// Stores an instance of the charmil factory
	cmdFactory *factory.Factory

	// Stores an instance of the charmil config handler
	h *config.CfgHandler

	// Initializes a zero-valued struct and stores its address
	opts = &options{}

	// Stores embedded contents of all the locales files
	//go:embed locales/*
	defaultLocales embed.FS

	// Stores an instance of the charmil localizer
	localizer localize.Localizer
)

func init() {
	// Links the specified local config file path and current config
	// struct pointer to a new charmil config handler instance and returns it
	h = config.NewHandler(cfgFilePath, opts)

	// Loads config values from the local config file
	err := h.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Initializes localizer config and adds it's value to the CLI config
	opts.LocConfig = localize.Config{
		Language: &language.English,
		Files:    defaultLocales,
		Format:   "yaml",
	}

	// Initializes the localizer by passing config
	localizer, err = localize.New(&opts.LocConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Creates a new factory instance with default settings
	cmdFactory = factory.Default(localizer, h)
}

func Root() *cobra.Command {
	// root command
	rootCmd := root.NewRootCommand(cmdFactory)
	rootCmd.InitDefaultHelpCmd()

	return rootCmd
}

func main() {
	cmd := Root()

	update.CheckForUpdate(context.Background(), cmdFactory.Logger, cmdFactory.Localizer)

	if err := doc.GenMarkdownTree(cmd, "./docs/commands"); err != nil {
		cmdFactory.Logger.Errorln(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	// Writes the current config into the local config file
	if err := h.Save(); err != nil {
		cmdFactory.Logger.Errorln(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		cmdFactory.Logger.Errorln(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}
}
