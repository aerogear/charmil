package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type createOptions struct {
	// Add your option fields here

	name         string
	outputFormat string
	autoUse      bool
	interactive  bool
}

// NewCreateCommand creates a new command for creating instances.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &createOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("crud.cmd.create.use"),
		Short:   f.Localizer.LocalizeByID("crud.cmd.create.shortDescription"),
		Long:    f.Localizer.LocalizeByID("crud.cmd.create.longDescription"),
		Example: f.Localizer.LocalizeByID("crud.cmd.create.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(opts, f)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", f.Localizer.LocalizeByID("crud.cmd.flag.output.description"))
	cmd.Flags().BoolVar(&opts.autoUse, "use", true, f.Localizer.LocalizeByID("crud.cmd.create.flag.use.description"))

	return cmd
}

func runCreate(opts *createOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
