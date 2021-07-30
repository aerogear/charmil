package root

import (
	initialize "github.com/aerogear/charmil/cli/internal/cmd/init"
	"github.com/aerogear/charmil/cli/internal/factory"
	"github.com/spf13/cobra"
)

func NewRootCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:           f.Localizer.LocalizeByID("charmil.cmd.use"),
		Short:         f.Localizer.LocalizeByID("charmil.cmd.short"),
		Long:          f.Localizer.LocalizeByID("charmil.cmd.long"),
		Example:       f.Localizer.LocalizeByID("charmil.cmd.example"),
		SilenceErrors: true,
	}

	cmd.AddCommand(initialize.InitCommand(f))

	return cmd
}
