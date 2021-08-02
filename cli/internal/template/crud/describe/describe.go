// Describe file for {{.Singular}} instances

package describe

import (
	"github.com/aerogear/charmil/cli/internal/factory"
	"github.com/spf13/cobra"
)

type describeOptions struct {
	id           string
	outputFormat string

	// You can add more fields here according to your requirements
}

// NewDescribeCommand creates a new command for describing instances.
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &describeOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("{{.Singular}}.cmd.describe.use"),
		Short:   f.Localizer.LocalizeByID("{{.Singular}}.cmd.describe.shortDescription"),
		Long:    f.Localizer.LocalizeByID("{{.Singular}}.cmd.describe.longDescription"),
		Example: f.Localizer.LocalizeByID("{{.Singular}}.cmd.describe.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDescribe(opts, f)
		},
	}

	// Adds local flags
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "json", f.Localizer.LocalizeByID("{{.Singular}}.cmd.flag.output.description"))
	cmd.Flags().StringVar(&opts.id, "id", "", f.Localizer.LocalizeByID("{{.Singular}}.common.flag.id"))

	return cmd
}
