package date

import (
	"time"

	"github.com/aerogear/charmil/pkg/factory"
	"github.com/aerogear/charmil/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Logger func() (logging.Logger, error)
}

func DateCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Logger: f.Logger,
	}
	cmd := &cobra.Command{
		Use:          "date",
		Short:        "tell date",
		Example:      "$ host date",
		SilenceUsage: true,
		Run: func(cmd *cobra.Command, args []string) {
			logger, _ := opts.Logger()
			logger.Output("Date Time is", time.Now())
		},
	}

	return cmd
}
