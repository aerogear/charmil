// REST API exposed via the serve command.
package starter

import (
	"github.com/aerogear/charmil/pkg/cmd/factory"

	"github.com/spf13/cobra"
)

func NewServiceRegistryCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "starter",
		Short:   f.Localizer.MustLocalize("root.cmd.shortDescription"),
		Long:    f.Localizer.MustLocalize("root.cmd.longDescription"),
		Example: f.Localizer.MustLocalize("root.cmd.example"),
		Args:    cobra.MinimumNArgs(1),
	}

	return cmd
}
