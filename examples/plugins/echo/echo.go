package echo

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/localize"
	"github.com/aerogear/charmil/core/logging"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

// Options is a type to access factory functions
// User can limit the options to use comming from factory
type Options struct {
	Logger   logging.Logger
	Localize localize.Localizer
}

// Date Command
func EchoCommand() (*cobra.Command, error) {

	// Initialize localizer providing the language, locals and format of locals file
	loc, err := localize.InitLocalizer(localize.Config{Language: language.English,
		Path: "examples/plugins/echo/locales/en/en.yaml", Format: "yaml"})
	if err != nil {
		return nil, err
	}

	// Create new/default instance of factory
	newFactory := factory.Default(loc)
	opts := &Options{
		Logger:   newFactory.Logger,
		Localize: newFactory.Localizer,
	}

	// creating new echo command
	// using localizer to access default text by ID provided in locals
	cmd := &cobra.Command{
		Use:     opts.Localize.LocalizeByID("echo.cmd.use"),
		Short:   opts.Localize.LocalizeByID("echo.cmd.short"),
		Long:    opts.Localize.LocalizeByID("echo.cmd.long"),
		Example: opts.Localize.LocalizeByID("echo.cmd.example"),
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {

			// Using logger for output
			opts.Logger.Info(args[0])

			return nil
		},
	}

	return cmd, nil
}
