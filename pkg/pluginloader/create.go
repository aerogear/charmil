package pluginloader

import (
	"strconv"

	"github.com/spf13/cobra"
)

// create a new cobra command
func CreateCommand(cmdStruct *CommandConfig, execute func(cmd *cobra.Command, args []string)) *cobra.Command {
	cmd := &cobra.Command{
		Use:     cmdStruct.Name,
		Short:   cmdStruct.ShortDescription,
		Example: cmdStruct.Examples,
		Args:    cobra.ExactArgs(len(cmdStruct.Args)),
		Run: func(cmd *cobra.Command, args []string) {
			execute(cmd, args)
		},
	}
	addFlags(cmdStruct, cmd)
	return cmd
}

// add Flags to the command
func addFlags(cmdStruct *CommandConfig, cmd *cobra.Command) {
	if cmdStruct.Flags != nil && len(cmdStruct.Flags) > 0 {
		for _, flag := range cmdStruct.Flags {
			switch flag.Type {
			case "string":
				cmd.Flags().StringP(flag.Name, flag.Alias, flag.DefaultValue, flag.Description)
			case "bool":
				v, _ := strconv.ParseBool(flag.DefaultValue)
				cmd.Flags().BoolP(flag.Name, flag.Alias, v, flag.Description)
			case "int":
				v, _ := strconv.Atoi(flag.DefaultValue)
				cmd.Flags().IntP(flag.Name, flag.Alias, v, flag.Description)
			}
		}
	}
}
