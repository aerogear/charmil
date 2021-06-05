package main

import (
	"io/ioutil"
	"log"
	"plugin"

	"github.com/aerogear/charmil/pkg/cmd/charmil"
	"github.com/spf13/cobra"
)

var (
	subcommands = []*cobra.Command{charmil.InstallCmd}

	cmd = &cobra.Command{
		Use:          "charmil",
		Short:        "Commands for managing plugins",
		SilenceUsage: true,
	}
)

func main() {
	// err = doc.GenMarkdownTree(cmd, "./docs/commands")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	pluginNames, err := ioutil.ReadDir("plugins/shared_objects/")
	if err != nil {
		log.Fatal(err)
	}
	for _, pluginName := range pluginNames {
		filePath := "plugins/shared_objects/" + pluginName.Name()

		p, err := plugin.Open(filePath)
		if err != nil {
			log.Fatal(err)
		}
		sym, err := p.Lookup("RootCmd")
		if err != nil {
			log.Fatal(err)
		}

		cmd.AddCommand(*sym.(**cobra.Command))
	}
}
