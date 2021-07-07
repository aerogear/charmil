package version

import (
	"fmt"

	"github.com/aerogear/charmil/internal/build"
	"github.com/aerogear/charmil/pkg/cmd/factory"
	"github.com/aerogear/charmil/pkg/iostreams"
	"github.com/aerogear/charmil/pkg/localize"
	"github.com/aerogear/charmil/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	IO        *iostreams.IOStreams
	Logger    func() (logging.Logger, error)
	localizer localize.Localizer
}

func NewVersionCmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:        f.IOStreams,
		Logger:    f.Logger,
		localizer: f.Localizer,
	}

	cmd := &cobra.Command{
		Use:    opts.localizer.MustLocalize("version.cmd.use"),
		Short:  opts.localizer.MustLocalize("version.cmd.shortDescription"),
		Hidden: true,
		Args:   cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCmd(opts)
		},
	}

	return cmd
}

func runCmd(opts *Options) (err error) {
	fmt.Fprintln(opts.IO.Out, opts.localizer.MustLocalize("version.cmd.outputText", localize.NewEntry("Version", build.Version)))
	return nil
}
