package main

import (
	"context"
	"embed"
	"fmt"
	"os"

	"github.com/aerogear/charmil/core/build"
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/localize"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"

	"github.com/aerogear/charmil/starter-1/pkg/cmd/root"
)

//go:embed locales
var defaultPath embed.FS

func abc() (*cobra.Command, *factory.Factory) {
	// init localizer
	localizer, err := localize.New(&localize.Config{
		Files:    defaultPath,
		Language: &language.English,
		Format:   "yaml",
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	buildVersion := build.Version
	// init factory
	cmdFactory := factory.Default(localizer)

	// root command
	rootCmd := root.NewRootCommand(cmdFactory, buildVersion)
	rootCmd.InitDefaultHelpCmd()

	return rootCmd, cmdFactory
}

func main() {
	cmd, cmdFactory := abc()

	cmd.Execute()
	build.CheckForUpdate(context.Background(), cmdFactory.Logger, cmdFactory.Localizer)
}
