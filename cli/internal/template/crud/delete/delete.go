// Delete file for {{.Singular}} instances

package delete

import (
	"github.com/aerogear/charmil/cli/internal/factory"
	"github.com/spf13/cobra"
)

type deleteOptions struct {
	id    string
	force bool

	// You can add more fields here according to your requirements
}

// GetDeleteCommand returns a new command for deleting {{.Singular}} instances.
func GetDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &deleteOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("{{.Singular}}.cmd.delete.use"),
		Short:   f.Localizer.LocalizeByID("{{.Singular}}.cmd.delete.shortDescription"),
		Long:    f.Localizer.LocalizeByID("{{.Singular}}.cmd.delete.longDescription"),
		Example: f.Localizer.LocalizeByID("{{.Singular}}.cmd.delete.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDelete(opts, f)
		},
	}

	// Adds local flags
	cmd.Flags().StringVar(&opts.id, "id", "", f.Localizer.LocalizeByID("{{.Singular}}.common.flag.id"))
	cmd.Flags().BoolVarP(&opts.force, "yes", "y", false, f.Localizer.LocalizeByID("{{.Singular}}.common.flag.yes"))

	return cmd
}
