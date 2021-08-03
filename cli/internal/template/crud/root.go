package crud

var RootTemplate = []byte(`package {{.Singular}}

import (
	"github.com/aerogear/charmil/cli/internal/factory"
	"github.com/spf13/cobra"
)

func NewCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("{{.Singular}}.cmd.use"),
		Short:   f.Localizer.LocalizeByID("{{.Singular}}.cmd.shortDescription"),
		Long:    f.Localizer.LocalizeByID("{{.Singular}}.cmd.longDescription"),
		Example: f.Localizer.LocalizeByID("{{.Singular}}.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	// Add sub-commands
	cmd.AddCommand(
		create.GetCreateCommand(f),
		describe.GetDescribeCommand(f),
		delete.GetDeleteCommand(f),
		list.GetListCommand(f),
		use.GetUseCommand(f),
	)

	return cmd
}`)
