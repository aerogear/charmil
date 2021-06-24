package date

import (
	"time"

	"github.com/aerogear/charmil/pkg/factory"
	"github.com/aerogear/charmil/pkg/localize"
	"github.com/aerogear/charmil/pkg/logging"
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
func DateCommand() (*cobra.Command, error) {

	// Initialize localizer providing the language, locals and format of locals file
	loc, err := localize.InitLocalizer(localize.Config{Language: language.English, Path: "examples/plugins/date/locals/en/en.yaml", Format: "yaml"})
	if err != nil {
		return nil, err
	}

	// Create new/default instance of factory
	newFactory := factory.Default(loc)
	opts := &Options{
		Logger:   newFactory.Logger,
		Localize: newFactory.Localizer,
	}

	// creating new command
	// using localizer to access default text by ID provided in locals
	cmd := &cobra.Command{
		Use:          opts.Localize.LocalizeByID("date.cmd.use"),
		Short:        opts.Localize.LocalizeByID("date.cmd.short"),
		Long:         opts.Localize.LocalizeByID("date.cmd.long"),
		Example:      opts.Localize.LocalizeByID("date.cmd.example"),
		Args:         cobra.ExactArgs(0),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {

			// Using logger for output
			opts.Logger.Infof("Date Time is %s", time.Now())

			return nil
		},
	}

	return cmd, nil
}
