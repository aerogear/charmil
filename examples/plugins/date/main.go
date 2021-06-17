package date

import (
	"fmt"
	"time"

	"github.com/aerogear/charmil/pkg/factory"
	"github.com/aerogear/charmil/pkg/localize"
	"github.com/aerogear/charmil/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Logger   func() (logging.Logger, error)
	Localize localize.Localizer
}

func DateCommand(f *factory.Factory) *cobra.Command {

	opts := &Options{
		Logger:   f.Logger,
		Localize: f.Localizer,
	}

	cmd := &cobra.Command{
		Use:          opts.Localize.LocalizeByID("date.use"),
		Short:        "tell date",
		Example:      "$ host date",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			logger, _ := opts.Logger()
			fmt.Println(opts.Localize.LocalizeByID("yo"))
			logger.Output("Date Time is", time.Now())
		},
	}

	return cmd
}
