package zsh

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

func NewCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:                   f.Localizer.LocalizeByID("completion.zsh.cmd.use"),
		Short:                 f.Localizer.LocalizeByID("completion.zsh.cmd.shortDescription"),
		Long:                  f.Localizer.LocalizeByID("completion.zsh.cmd.longDescription"),
		DisableFlagsInUseLine: true,
		Args:                  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Parent().Parent().GenZshCompletion(f.IOStreams.Out)
		},
	}

	return cmd
}
