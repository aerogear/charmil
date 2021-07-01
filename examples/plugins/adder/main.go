package adder

import (
	"os"
	"strconv"

	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/localize"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

func AdderCommand() (*cobra.Command, error) {
	locConfig := localize.Config{
		Language: language.English,
		Path:     "examples/plugins/adder/locales/en/adder.en.yaml",
		Format:   "yaml",
	}

	loc, err := localize.InitLocalizer(locConfig)
	if err != nil {
		return nil, err
	}

	f := factory.Default(loc)

	adderCmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("adder.cmd.use"),
		Short:   f.Localizer.LocalizeByID("adder.cmd.short"),
		Example: f.Localizer.LocalizeByID("adder.cmd.example"),
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			result := 0

			for _, arg := range args {
				n, err := strconv.Atoi(arg)
				if err != nil {
					f.Logger.Errorln(err)
					os.Exit(1)
				}
				result += n
			}

			f.Logger.Infof("The result is: %d\n", result)
		},
	}

	return adderCmd, nil
}
