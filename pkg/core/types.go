package core

import "github.com/spf13/cobra"

type Plugin interface {
	CreateRootCmd( /*TODO: Define config struct to add as arg here*/ ) *cobra.Command
}
