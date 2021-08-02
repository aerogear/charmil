// List file for {{.Singular}} instances

package list

import (
	"github.com/aerogear/charmil/cli/internal/factory"
	"github.com/spf13/cobra"
)

type listOptions struct {
	outputFormat string
	page         int32
	limit        int32
	search       string

	// You can add more fields here according to your requirements
}

// NewListCommand creates a new command for listing instances.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &listOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("{{.Singular}}.cmd.list.use"),
		Short:   f.Localizer.LocalizeByID("{{.Singular}}.cmd.list.shortDescription"),
		Long:    f.Localizer.LocalizeByID("{{.Singular}}.cmd.list.longDescription"),
		Example: f.Localizer.LocalizeByID("{{.Singular}}.cmd.list.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(opts, f)
		},
	}

	// Adds local flags
	cmd.Flags().StringVarP(&opts.outputFormat, "output", "o", "", f.Localizer.LocalizeByID("{{.Singular}}.cmd.flag.output.description"))
	cmd.Flags().Int32VarP(&opts.page, "page", "", 1, f.Localizer.LocalizeByID("{{.Singular}}.list.flag.page"))
	cmd.Flags().Int32VarP(&opts.limit, "limit", "", 100, f.Localizer.LocalizeByID("{{.Singular}}.list.flag.limit"))
	cmd.Flags().StringVarP(&opts.search, "search", "", "", f.Localizer.LocalizeByID("{{.Singular}}.list.flag.search"))

	return cmd
}
