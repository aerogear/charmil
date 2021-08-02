// Use file for {{.Singular}} instances

package use

import (
	"github.com/aerogear/charmil/cli/internal/factory"
	"github.com/spf13/cobra"
)

type useOptions struct {
	id string

	// You can add more fields here according to your requirements
}

// NewUseCommand creates a new command for using instances.
func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &useOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("{{.Singular}}.cmd.use.use"),
		Short:   f.Localizer.LocalizeByID("{{.Singular}}.cmd.use.shortDescription"),
		Long:    f.Localizer.LocalizeByID("{{.Singular}}.cmd.use.longDescription"),
		Example: f.Localizer.LocalizeByID("{{.Singular}}.cmd.use.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUse(opts, f)
		},
	}

	// Adds local flags
	cmd.Flags().StringVar(&opts.id, "id", "", f.Localizer.LocalizeByID("{{.Singular}}.use.flag.id"))

	return cmd
}
