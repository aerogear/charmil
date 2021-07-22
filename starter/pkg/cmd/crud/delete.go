package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type deleteOptions struct {
	// Add your option fields here

	id    string
	name  string
	force bool
}

// NewDeleteCommand creates a new command for deleting instances.
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &deleteOptions{}

	cmd := &cobra.Command{
		Use:     "delete",
		Short:   f.Localizer.LocalizeByID("crud.cmd.delete.shortDescription"),
		Long:    f.Localizer.LocalizeByID("crud.cmd.delete.longDescription"),
		Example: f.Localizer.LocalizeByID("crud.cmd.delete.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDelete(opts, f)
		},
	}

	cmd.Flags().StringVar(&opts.id, "id", "", f.Localizer.LocalizeByID("crud.common.flag.id"))
	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, f.Localizer.LocalizeByID("crud.common.flag.yes"))

	return cmd
}

func runDelete(opts *deleteOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
