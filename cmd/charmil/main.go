package main

import (
	"log"

	"github.com/aerogear/charmil/pkg/cmd/charmil"
	"github.com/aerogear/charmil/pkg/core"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

const cliKey = "plugin"

var subcommands = []*cobra.Command{charmil.InstallCmd, charmil.InstalledCmd, charmil.ListCmd, charmil.RemoveCmd, charmil.ActivateCmd, charmil.DeactivateCmd}

func main() {
	cmd := &cobra.Command{
		Use:          "charmil",
		Short:        "Commands for managing plugins",
		SilenceUsage: true,
	}

	charmilCLI := core.GetCLI(cliKey)
	err := charmilCLI.AddCommands(cmd, subcommands...)
	if err != nil {
		log.Fatal(err)
	}

	err = doc.GenMarkdownTree(cmd, "./docs/commands")
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
