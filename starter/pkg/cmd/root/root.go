package root

import (
	"flag"

	"github.com/aerogear/charmil/starter/pkg/cmd/starter"

	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewRootCommand(f *factory.Factory) *cobra.Command {

	cmd := &cobra.Command{
		SilenceUsage:  true,
		SilenceErrors: true,
		Use:           f.Localizer.LocalizeByID("root.cmd.use"),
		Short:         f.Localizer.LocalizeByID("root.cmd.shortDescription"),
		Long:          f.Localizer.LocalizeByID("root.cmd.longDescription"),
		Example:       f.Localizer.LocalizeByID("root.cmd.example"),
	}
	fs := cmd.PersistentFlags()

	// this flag comes out of the box, but has its own basic usage text, so this overrides that
	var help bool

	fs.BoolVarP(&help, "help", "h", false, f.Localizer.LocalizeByID("root.cmd.flag.help.description"))

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)

	cmd.AddCommand(starter.NewServiceRegistryCommand(f))

	// Uncomment the following snippet to add a plugin to this CLI:
	/*
		pluginRootCmd, err := pluginAlias.RootCommand(f)
		if err != nil {
			f.Logger.Errorln(f.IOStreams.ErrOut, err)
			os.Exit(1)
		}
		cmd.AddCommand(pluginRootCmd)
	*/

	return cmd
}
