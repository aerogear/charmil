package core

import (
	"github.com/aerogear/charmil/pkg/factory"
	"github.com/spf13/cobra"
)

func GetPluginRootCmd(cmd CLI, f *factory.Factory) *cobra.Command {
	return cmd.NewRootCmd(f)
}
