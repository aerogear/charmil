package fish

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

func NewCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   f.Localizer.LocalizeByID("completion.fish.cmd.use"),
		Short:                 f.Localizer.LocalizeByID("completion.fish.cmd.shortDescription"),
		Long:                  f.Localizer.LocalizeByID("completion.fish.cmd.longDescription"),
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Parent().GenFishCompletion(f.IOStreams.Out, false)
		},
	}

	return cmd
}
