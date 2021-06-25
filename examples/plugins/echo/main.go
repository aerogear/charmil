package echo

import (
	"errors"
	"strings"

	"github.com/aerogear/charmil/pkg/factory"
	"github.com/aerogear/charmil/pkg/localize"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

func EchoCommand() (*cobra.Command, error) {
	var (
		iter                 int
		uppercase, lowercase bool
	)

	locConfig := localize.Config{
		Language: language.English,
		Path:     "examples/plugins/echo/locals/en/echo.en.yaml",
		Format:   "yaml",
	}

	loc, err := localize.InitLocalizer(locConfig)
	if err != nil {
		return nil, err
	}

	f := factory.Default(loc)

	echoCmd := &cobra.Command{
		Use:          f.Localizer.LocalizeByID("echo.cmd.use"),
		Short:        f.Localizer.LocalizeByID("echo.cmd.short"),
		Example:      f.Localizer.LocalizeByID("echo.cmd.example"),
		SilenceUsage: true,
		Args:         cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			message := strings.Join(args, " ")

			if uppercase {
				message = strings.ToUpper(message)
			} else if lowercase {
				message = strings.ToLower(message)
			}

			if iter == 0 {
				return errors.New("Number of iterations cannot be 0")
			}
			for i := 0; i < iter; i++ {
				f.Logger.Infof("Echo: %s\n", message)
			}

			return nil
		},
	}

	echoCmd.Flags().IntVarP(&iter, "repeat", "r", 1, f.Localizer.LocalizeByID("echo.cmd.flag.repeat.description"))
	echoCmd.Flags().BoolVarP(&uppercase, "uppercase", "u", false, f.Localizer.LocalizeByID("echo.cmd.flag.uppercase.description"))
	echoCmd.Flags().BoolVarP(&lowercase, "lowercase", "l", false, f.Localizer.LocalizeByID("echo.cmd.flag.lowercase.description"))

	return echoCmd, nil
}
