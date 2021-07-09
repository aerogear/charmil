package init

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

func InitCommand(f *factory.Factory, version string) *cobra.Command {

	cmd := &cobra.Command{
		Use:           f.Localizer.LocalizeByID("init.cmd.use"),
		Short:         f.Localizer.LocalizeByID("init.cmd.short"),
		Long:          f.Localizer.LocalizeByID("init.cmd.long"),
		Example:       f.Localizer.LocalizeByID("init.cmd.example"),
		SilenceErrors: true,
		Run:           func(cmd *cobra.Command, args []string) {},
	}

	cmd.Version = version

	return cmd
}
