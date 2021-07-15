package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/localize"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"

	"github.com/aerogear/charmil/starter-1/pkg/cmd/root"
)

//go:embed locales
var defaultPath embed.FS

func abc() *cobra.Command {
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

	// init factory
	cmdFactory := factory.Default(localizer)

	// root command
	rootCmd := root.NewRootCommand(cmdFactory)
	rootCmd.InitDefaultHelpCmd()

	return rootCmd
}

func main() {
	cmd := abc()

	if err := cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
