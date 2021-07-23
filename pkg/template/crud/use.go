// Use file for {{.Singular}} instances

package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type useOptions struct {
	// Add your option fields here

}

// NewUseCommand creates a new command for using instances.
func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &useOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("crud.cmd.use.use"),
		Short:   f.Localizer.LocalizeByID("crud.cmd.use.shortDescription"),
		Long:    f.Localizer.LocalizeByID("crud.cmd.use.longDescription"),
		Example: f.Localizer.LocalizeByID("crud.cmd.use.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUse(opts, f)
		},
	}

	return cmd
}

func runUse(opts *useOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
