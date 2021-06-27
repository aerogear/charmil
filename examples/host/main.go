package main

import (
	"fmt"
	"log"

	"github.com/aerogear/charmil/examples/plugins/date"
	"github.com/aerogear/charmil/examples/plugins/echo"
	"github.com/aerogear/charmil/pkg/preview"
	"github.com/spf13/cobra"
)

func main() {
	// Root command of the host CLI
	cmd := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
		Run:          preview.UpdateFlagValue,
	}

	// Adds the `dev_preview` flag to root command
	preview.InitFlag(cmd)

	// Stores the root command of the `date` plugin
	dateCmd, err := date.DateCommand()
	if err != nil {
		fmt.Println(err)
	}

	// Stores the root command of the `echo` plugin
	echoCmd, err := echo.EchoCommand()
	if err != nil {
		fmt.Println(err)
	}

	// Adds `echoCmd` as a command available under dev preview
	preview.AddCommands(echoCmd)

	// Adds the root command of plugins as child commands
	cmd.AddCommand(dateCmd)
	cmd.AddCommand(echoCmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
