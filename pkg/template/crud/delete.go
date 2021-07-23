// Delete file for {{.Singular}} instances

package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type deleteOptions struct {
	// Add your option fields here

}

// NewDeleteCommand creates a new command for deleting instances.
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &deleteOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("crud.cmd.delete.use"),
		Short:   f.Localizer.LocalizeByID("crud.cmd.delete.shortDescription"),
		Long:    f.Localizer.LocalizeByID("crud.cmd.delete.longDescription"),
		Example: f.Localizer.LocalizeByID("crud.cmd.delete.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDelete(opts, f)
		},
	}

	return cmd
}

func runDelete(opts *deleteOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
