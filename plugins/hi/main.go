package hi

import (
	"fmt"

	"github.com/aerogear/charmil/pkg/pluginloader"
	"github.com/spf13/cobra"
)

var HiCommand = pluginloader.CreateCommand(&pluginloader.CommandConfig{
	Name:             "hi",
	Args:             []string{"name"},
	Flags:            []pluginloader.FlagConfig{},
	ShortDescription: "say hi",
	Examples:         "$ host hi ankit",
}, func(cmd1 *cobra.Command, args []string) {
	for _, ival := range args {
		fmt.Printf("You're welcome " + ival)
	}
})
