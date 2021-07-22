package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type describeOptions struct {
	// Add your option fields here

	id           string
	name         string
	outputFormat string
}

// NewDescribeCommand creates a new command for describing instances.
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &describeOptions{}

	cmd := &cobra.Command{
		Use:     "describe",
		Short:   f.Localizer.LocalizeByID("crud.cmd.describe.shortDescription"),
		Long:    f.Localizer.LocalizeByID("crud.cmd.describe.longDescription"),
		Example: f.Localizer.LocalizeByID("crud.cmd.describe.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDescribe(opts, f)
		},
	}

	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", f.Localizer.LocalizeByID("crud.cmd.flag.output.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", f.Localizer.LocalizeByID("crud.common.flag.id"))

	return cmd
}

func runDescribe(opts *describeOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
