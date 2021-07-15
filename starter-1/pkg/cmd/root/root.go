package root

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"

	cliversion "github.com/aerogear/charmil/starter-1/pkg/cmd/version"
)

func NewRootCommand(f *factory.Factory) *cobra.Command {

	cmd := &cobra.Command{
		Use:           f.Localizer.LocalizeByID("root.cmd.use"),
		Short:         f.Localizer.LocalizeByID("root.cmd.short"),
		Long:          f.Localizer.LocalizeByID("root.cmd.long"),
		Example:       f.Localizer.LocalizeByID("root.cmd.example"),
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.AddCommand(cliversion.NewVersionCmd(f))

	return cmd
}
