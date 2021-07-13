package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	c "github.com/aerogear/charmil/core/config"
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/localize"
	"github.com/aerogear/charmil/starter/internal/build"
	"github.com/aerogear/charmil/starter/pkg/cmd/root"
	"github.com/redhat-developer/app-services-cli/pkg/doc"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

var generateDocs = os.Getenv("GENERATE_DOCS") == "true"

// Defines the configuration keys of the host CLI.
//
// CONSTRAINT: All fields of the config struct need to be exportable
type config struct {
	LocConfig localize.Config
	Plugins   map[string]interface{}
}

// Stores the path of the local config file
const cfgFilePath = "./config.json"

var (
	// Stores an instance of the charmil factory
	cmdFactory *factory.Factory

	// Stores an instance of the charmil config handler
	h *c.CfgHandler

	// Initializes a zero-valued struct and stores its address
	cfg = &config{}

	//go:embed locales/*
	defaultLocales embed.FS

	localizer localize.Localizer

	buildVersion string
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

	// TODO: Execute this only when the local config file doesn't have a LocConfig value already present
	cfg.LocConfig = localize.Config{
		Language: &language.English,
		Files:    defaultLocales,
		Format:   "yaml",
	}

	localizer, err := localize.New(&cfg.LocConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Creates a new factory instance with default settings
	cmdFactory = factory.Default(localizer, h)

	buildVersion = build.Version
}

func main() {
	rootCmd := root.NewRootCommand(cmdFactory, buildVersion)

	rootCmd.InitDefaultHelpCmd()

	// Writes the current config into the local config file
	err := h.Save()
	if err != nil {
		cmdFactory.Logger.Errorln(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	if err = rootCmd.Execute(); err != nil {
		cmdFactory.Logger.Errorln(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	if generateDocs {
		generateDocumentation(rootCmd)
	}
}

/**
* Generates documentation files
 */
func generateDocumentation(rootCommand *cobra.Command) {
	fmt.Fprint(os.Stderr, "Generating docs.\n\n")
	filePrepender := func(filename string) string {
		return ""
	}

	rootCommand.DisableAutoGenTag = true

	linkHandler := func(s string) string { return s }

	err := doc.GenAsciidocTreeCustom(rootCommand, "./docs/commands", filePrepender, linkHandler)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
