package echo

import (
	"errors"
	"strings"

	"github.com/aerogear/charmil/pkg/factory"
	"github.com/aerogear/charmil/pkg/localize"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

// EchoCommand returns echo plugin's root command struct.
// This will be added to the host CLI as an extension.
func EchoCommand() (*cobra.Command, error) {
	var (
		// Stores the number of iterations specified through the `repeat` flag.
		iter int

		// Holds a `true` value if the uppercase flag is set.
		uppercase bool

		// Holds a `true` value if the lowercase flag is set.
		lowercase bool
	)

	// Stores the config for localizer
	locConfig := localize.Config{
		Language: language.English,
		Path:     "examples/plugins/echo/locals/en/echo.en.yaml",
		Format:   "yaml",
	}

	// Initializes the localizer by passing config
	loc, err := localize.InitLocalizer(locConfig)
	if err != nil {
		return nil, err
	}

	// Creates a new factory instance using default configuration
	f := factory.Default(loc)

	// Stores the root command of plugin
	echoCmd := &cobra.Command{
		Use:          f.Localizer.LocalizeByID("echo.cmd.use"),
		Short:        f.Localizer.LocalizeByID("echo.cmd.short"),
		Example:      f.Localizer.LocalizeByID("echo.cmd.example"),
		SilenceUsage: true,
		Args:         cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			message := strings.Join(args, " ")

			// Handles the case when any case converter flag is set
			if uppercase {
				message = strings.ToUpper(message)
			} else if lowercase {
				message = strings.ToLower(message)
			}

			// Handles the base case
			if iter < 1 {
				return errors.New("Number of iterations cannot be less than 1")
			}

			// Prints the specified message to stdout
			for i := 0; i < iter; i++ {
				f.Logger.Infof("Echo: %s\n", message)
			}

			return nil
		},
	}

	// Adds local flags to the root command
	echoCmd.Flags().IntVarP(&iter, "repeat", "r", 1, f.Localizer.LocalizeByID("echo.cmd.flag.repeat.description"))
	echoCmd.Flags().BoolVarP(&uppercase, "uppercase", "u", false, f.Localizer.LocalizeByID("echo.cmd.flag.uppercase.description"))
	echoCmd.Flags().BoolVarP(&lowercase, "lowercase", "l", false, f.Localizer.LocalizeByID("echo.cmd.flag.lowercase.description"))

	return echoCmd, nil
}
