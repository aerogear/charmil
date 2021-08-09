package root

import (
	"fmt"

	"github.com/aerogear/charmil/cli/internal/cmd/add"
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

	addCmd, err := add.AddCommand(f)
	if err != nil {
		fmt.Println(err)
	}
	cmd.AddCommand(addCmd)

	return cmd
}
