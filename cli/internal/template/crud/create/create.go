// Create file for {{.Singular}} instances

package create

import (
	"github.com/aerogear/charmil/cli/internal/factory"
	"github.com/spf13/cobra"
)

type createOptions struct {
	outputFormat string
	autoUse      bool

	// You can add more fields here according to your requirements
}

// NewCreateCommand creates a new command for creating instances.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &createOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("{{.Singular}}.cmd.create.use"),
		Short:   f.Localizer.LocalizeByID("{{.Singular}}.cmd.create.shortDescription"),
		Long:    f.Localizer.LocalizeByID("{{.Singular}}.cmd.create.longDescription"),
		Example: f.Localizer.LocalizeByID("{{.Singular}}.cmd.create.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(opts, f)
		},
	}

	// Adds local flags
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", f.Localizer.LocalizeByID("{{.Singular}}.cmd.flag.output.description"))
	cmd.Flags().BoolVar(&opts.autoUse, "use", true, f.Localizer.LocalizeByID("{{.Singular}}.cmd.create.flag.use.description"))

	return cmd
}
