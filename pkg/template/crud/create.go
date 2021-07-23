// Create file for {{.Singular}} instances

package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type createOptions struct {
	// Add your option fields here
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

	return cmd
}

func runCreate(opts *createOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
