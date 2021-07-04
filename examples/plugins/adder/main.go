package adder

import (
	"fmt"
	"log"
	"strconv"

	c "github.com/aerogear/charmil/core/config"
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/localize"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

type config struct {
	Key5 string
	Key6 string
	Key7 string
	Key8 string
}

var cfg = &config{}

// AdderCommand returns the root command of plugin.
// This will be added to the host CLI as an extension.
func AdderCommand(cfile c.CfgFile) (*cobra.Command, error) {
	// Stores a new instance of the charmil config handler
	// Links the handler instance to a local config file
	h := c.New(cfile, cfg)

	// // Loads config values from the local config file
	// err := h.LoadPluginCfg()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Sets dummy values into config
	cfg.Key5 = "val5"
	cfg.Key6 = "val6"
	cfg.Key7 = "val7"
	cfg.Key8 = "val8"

	// Stores the config for localizer
	locConfig := localize.Config{
		Language: language.English,
		Path:     "examples/plugins/adder/locales/en/adder.en.yaml",
		Format:   "yaml",
	}

	// Initializes the localizer by passing config
	loc, err := localize.InitLocalizer(locConfig)
	if err != nil {
		return nil, err
	}

	//Stores a new factory instance with default settings
	f := factory.Default(loc)

	// Stores the root command of plugin
	adderCmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("adder.cmd.use"),
		Short:   f.Localizer.LocalizeByID("adder.cmd.short"),
		Example: f.Localizer.LocalizeByID("adder.cmd.example"),
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

			f.Logger.Infoln(f.Localizer.LocalizeByID("adder.cmd.resultMessage"), result)

			return nil
		},
	}
	fmt.Println("Adder Plugin Config:", *cfg)

	// Writes the current config into the local config file
	err = h.SavePluginCfg()
	if err != nil {
		log.Fatal(err)
	}

	// Returns the root command of plugin along with the plugin config map
	return adderCmd, nil
}
