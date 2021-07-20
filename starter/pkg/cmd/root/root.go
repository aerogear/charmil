package root

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/starter/pkg/cmd/completion"
	"github.com/spf13/cobra"
)

func NewRootCommand(f *factory.Factory) *cobra.Command {

	cmd := &cobra.Command{
		Use:           f.Localizer.LocalizeByID("root.cmd.use"),
		Short:         f.Localizer.LocalizeByID("root.cmd.short"),
		Long:          f.Localizer.LocalizeByID("root.cmd.long"),
		Example:       f.Localizer.LocalizeByID("root.cmd.example"),
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	cmd.AddCommand(completion.NewCompletionCommand(f))

	return cmd
}
