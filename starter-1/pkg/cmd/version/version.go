package version

import (
	"fmt"

	"github.com/aerogear/charmil/core/build"
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/iostreams"
	"github.com/aerogear/charmil/core/localize"
	"github.com/aerogear/charmil/core/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	IO        *iostreams.IOStreams
	Logger    logging.Logger
	localizer localize.Localizer
}

func NewVersionCmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:        f.IOStreams,
		Logger:    f.Logger,
		localizer: f.Localizer,
	}

	cmd := &cobra.Command{
		Use:    opts.localizer.LocalizeByID("version.cmd.use"),
		Short:  opts.localizer.LocalizeByID("version.cmd.short"),
		Hidden: true,
		Args:   cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCmd(opts)
		},
	}

	return cmd
}

func runCmd(opts *Options) (err error) {
	fmt.Fprintln(opts.IO.Out, opts.localizer.LocalizeByID("version.cmd.outputText", localize.NewEntry("Version", build.Version)))
	return nil
}
