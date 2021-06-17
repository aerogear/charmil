package date

import (
	"time"

	"github.com/aerogear/charmil/pkg/factory"
	"github.com/aerogear/charmil/pkg/localize"
	"github.com/aerogear/charmil/pkg/logging"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

type Options struct {
	Logger   func() (logging.Logger, error)
	Localize localize.Localizer
}

func DateCommand() *cobra.Command {

	loc, err := localize.InitLocalizer(localize.Config{Language: language.English, Path: "examples/plugins/date/locals/en/en.yaml", Format: "e"})
	if err != nil {
		panic(err)
	}

	newFactory := factory.Default(loc)

	opts := &Options{
		Logger:   newFactory.Logger,
		Localize: newFactory.Localizer,
	}

	cmd := &cobra.Command{
		Use:          opts.Localize.LocalizeByID("date.cmd.use"),
		Short:        opts.Localize.LocalizeByID("date.cmd.short"),
		Example:      opts.Localize.LocalizeByID("date.cmd.example"),
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			logger, _ := opts.Logger()
			logger.Output("Date Time is", time.Now())
		},
	}

	return cmd
}
