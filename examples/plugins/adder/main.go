package adder

import (
	"strconv"

	"github.com/aerogear/charmil/core/config"
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/localize"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

// AdderCommand returns the root command of plugin.
// This will be added to the host CLI as an extension.
func AdderCommand() (*cobra.Command, map[string]interface{}, error) {
	// Stores a new instance of the charmil config handler
	h := config.New()

	// Sets dummy values into config map
	h.SetValue("key5", "val5")
	h.SetValue("key6", "val6")
	h.SetValue("key7", "val7")
	h.SetValue("key8", "val8")

	// Stores the config for localizer
	locConfig := localize.Config{
		Language: language.English,
		Path:     "examples/plugins/adder/locales/en/adder.en.yaml",
		Format:   "yaml",
	}

	// Initializes the localizer by passing config
	loc, err := localize.InitLocalizer(locConfig)
	if err != nil {
		return nil, nil, err
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

	// Returns the root command of plugin along with the plugin config map
	return adderCmd, h.GetAllSettings(), nil
}
