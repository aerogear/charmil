package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/aerogear/charmil/starter/internal/build"
	"golang.org/x/text/language"

	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/localize"
	"github.com/aerogear/charmil/starter/pkg/cmd/root"
)

// var generateDocs = os.Getenv("GENERATE_DOCS") == "true"

var (
	//go:embed locales/*
	defaultLocales embed.FS
)

func init() {

}

func main() {
	LocConfig := localize.Config{
		Language: &language.English,
		Files:    defaultLocales,
		Format:   "yaml",
	}

	localizer, err := localize.New(&LocConfig)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	buildVersion := build.Version
	cmdFactory := factory.Default(localizer, nil)

	if err != nil {
		fmt.Println(cmdFactory.IOStreams.ErrOut, err)
		os.Exit(1)
	}

	// initConfig(cmdFactory)

	rootCmd := root.NewRootCommand(cmdFactory, buildVersion)

	rootCmd.InitDefaultHelpCmd()

	// if generateDocs {
	// 	generateDocumentation(rootCmd)
	// 	os.Exit(0)
	// }

	err = rootCmd.Execute()

	if err == nil {
		return
	}
}

/**
* Generates documentation files
 */
// func generateDocumentation(rootCommand *cobra.Command) {
// 	fmt.Fprint(os.Stderr, "Generating docs.\n\n")
// 	filePrepender := func(filename string) string {
// 		return ""
// 	}

// 	rootCommand.DisableAutoGenTag = true

// 	linkHandler := func(s string) string { return s }

// 	err := doc.GenAsciidocTreeCustom(rootCommand, "./docs/commands", filePrepender, linkHandler)
// 	if err != nil {
// 		fmt.Fprintln(os.Stderr, err)
// 		os.Exit(1)
// 	}
// }

// func initConfig(f *factory.Factory) {
// 	cfgFile, err := f.Config.Load()

// 	if cfgFile != nil {
// 		return
// 	}
// 	if !os.IsNotExist(err) {
// 		fmt.Fprintln(f.IOStreams.ErrOut, err)
// 		os.Exit(1)
// 	}

// 	cfgFile = &config.Config{}
// 	if err := f.Config.Save(cfgFile); err != nil {
// 		fmt.Fprintln(f.IOStreams.ErrOut, err)
// 		os.Exit(1)
// 	}
// }
