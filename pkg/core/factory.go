package core

import (
	"log"

	"github.com/spf13/cobra"
)

const (
	hostKey   = "host"
	pluginKey = "plugin"
)

type CLI interface {
	AddCommands(cmd *cobra.Command, subcommands ...*cobra.Command) error
}

// cliConfig represents configuration settings common to both Host and Plugin CLI
type cliConfig struct {
	// This is similar to factory struct in: https://github.com/redhat-developer/app-services-cli/blob/main/pkg/cmd/factory/factory.go#L13-L24

	// Todo: Add fields
}

type host struct {
	*cliConfig
	// Todo: Add more fields
}
type plugin struct {
	*cliConfig
	// Todo: Add more fields
}

func (h host) AddCommands(cmd *cobra.Command, subcommands ...*cobra.Command) error {
	cmd.AddCommand(subcommands...)

	// Todo: Add config logic
	return nil
}

func (p plugin) AddCommands(cmd *cobra.Command, subcommands ...*cobra.Command) error {
	cmd.AddCommand(subcommands...)

	// Todo: Add config logic
	return nil
}

func GetCLI(cliType string) CLI {
	switch cliType {
	case hostKey:
		return &host{} // Todo: initialize struct with config values
	case pluginKey:
		return &plugin{} // Todo: initialize struct with config values
	default:
		log.Printf("Undefined CLI type")
		return nil
	}
}
