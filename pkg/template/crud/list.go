// List file for {{.Singular}} instances

package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type listOptions struct {
	// Add your option fields here

}

// NewListCommand creates a new command for listing instances.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &listOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("crud.cmd.list.use"),
		Short:   f.Localizer.LocalizeByID("crud.cmd.list.shortDescription"),
		Long:    f.Localizer.LocalizeByID("crud.cmd.list.longDescription"),
		Example: f.Localizer.LocalizeByID("crud.cmd.list.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(opts, f)
		},
	}

	return cmd
}

func runList(opts *listOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
