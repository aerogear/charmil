package echo

import (
	"embed"
	"log"

	c "github.com/aerogear/charmil/core/config"
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/localize"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

var (
	//go:embed locales
	defaultLocales embed.FS
)

// Defines the configuration keys of the plugin.
//
// CONSTRAINT: All fields of the config struct need to be exportable
type config struct {
	LocConfig localize.Config
}

// Initializes a zero-valued struct and stores its address
var cfg = &config{}

// EchoCommand returns the root command of plugin.
// This will be added to the host CLI as an extension.
func EchoCommand(f *factory.Factory) (*cobra.Command, error) {

	// Stores the config for localizer
	cfg.LocConfig = localize.Config{
		Language: &language.English,
		Files:    defaultLocales,
		Path:     "examples/plugins/echo/locales",
		Format:   "yaml",
	}

	// Initialize localizer providing the language, locals and format of locals file
	loc, err := localize.New(&cfg.LocConfig)
	if err != nil {
		return nil, err
	}

	// User can limit the options to use comming from factory
	opts := &factory.Factory{
		Logger:     f.Logger,
		Localizer:  loc,
		CfgHandler: f.CfgHandler,
	}

	// creating new echo command
	// using localizer to access default text by ID provided in locals
	cmd := &cobra.Command{
		Use:     opts.Localizer.LocalizeByID("echo.cmd.use"),
		Short:   opts.Localizer.LocalizeByID("echo.cmd.short"),
		Long:    opts.Localizer.LocalizeByID("echo.cmd.long"),
		Example: opts.Localizer.LocalizeByID("echo.cmd.example"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			// Using logger for output
			opts.Logger.Info(args[0])

			return nil
		},
	}

	// Merges the current plugin config into the host CLI config
	err = c.MergePluginCfg("echo", opts.CfgHandler, cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cmd, nil
}
