package main

import (
	"fmt"
	"log"

	"github.com/aerogear/charmil/examples/plugins/date"
	"github.com/aerogear/charmil/examples/plugins/echo"
	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}

	dateCmd, err := date.DateCommand()
	if err != nil {
		fmt.Println(err)
	}
	echoCmd, err := echo.EchoCommand()
	if err != nil {
		fmt.Println(err)
	}

	cmd.AddCommand(dateCmd)
	cmd.AddCommand(echoCmd)

	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
