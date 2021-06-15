package pluginA

import (
	"github.com/aerogear/charmil/pkg/factory"
	"github.com/spf13/cobra"
)

type myPlugin struct{}

func GetPlugin() *myPlugin {
	return &myPlugin{}
}

func (p *myPlugin) NewRootCmd(f *factory.Factory) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "pluginA",
		Short: "Plugin A",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger, err := f.Logger()
			if err != nil {
				return err
			}

			logger.Info("Root command of Plugin A called")
			return nil
		},
	}
	return rootCmd
}
