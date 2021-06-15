package core

import (
	"github.com/aerogear/charmil/pkg/factory"
	"github.com/spf13/cobra"
)

type CLI interface {
	NewRootCmd(f *factory.Factory) *cobra.Command
}
