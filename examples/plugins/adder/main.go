package adder

import (
	"embed"
	"log"
	"strconv"

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

// AdderCommand returns the root command of plugin.
// This will be added to the host CLI as an extension.
func AdderCommand(f *factory.Factory) (*cobra.Command, error) {

	// Stores the config for localizer
	cfg.LocConfig = localize.Config{
		Language: &language.English,
		Files:    defaultLocales,
		Path:     "examples/plugins/adder/locales",
		Format:   "yaml",
	}

	// Initializes the localizer by passing config
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

	// Stores the root command of plugin
	adderCmd := &cobra.Command{
		Use:     opts.Localizer.LocalizeByID("adder.cmd.use"),
		Short:   opts.Localizer.LocalizeByID("adder.cmd.short"),
		Example: opts.Localizer.LocalizeByID("adder.cmd.example"),
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			result := 0

			for _, arg := range args {
				n, err := strconv.Atoi(arg)
				if err != nil {
					return err
				}
				result += n
			}

			opts.Logger.Infoln(opts.Localizer.LocalizeByID("adder.cmd.resultMessage"), result)

			return nil
		},
	}

	// Merges the current plugin config into the host CLI config
	err = c.MergePluginCfg("adder", opts.CfgHandler, cfg)
	if err != nil {
		log.Fatal(err)
	}

	return adderCmd, nil
}
