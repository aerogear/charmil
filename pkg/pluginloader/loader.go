package pluginloader

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

// TODO:
// Unmarshal the YAML
// Extract the commands using plugin.go (command structure)

// Create a Cobra Command New
func createCommandNew(cmdStruct *CommandConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:     cmdStruct.Name,
		Short:   cmdStruct.ShortDescription,
		Example: cmdStruct.Examples,
		Args:    cobra.ExactArgs(len(cmdStruct.MapsTo.Args)),
		RunE: func(cmd *cobra.Command, args []string) error {

			// TODO: arguments are not typed
			// append subcommand before the []args
			args = append([]string{cmdStruct.MapsTo.Subcommand}, args...)

			return Execute(cmdStruct, cmdStruct.MapsTo.Name, args)
		},
	}
	addFlags(cmdStruct, cmd)
	return cmd
}

// Execute external process/command
func Execute(cmdStruct *CommandConfig, executablePath string, cmdArgs []string) error {

	// If command has flags
	flag.Parse()
	tail := flag.Args()
	f := tail[len(tail)-1]
	if strings.HasPrefix(f, "-") || strings.HasPrefix(f, "--") { // checking last element in the tail array
		for _, fl := range cmdStruct.Flags {
			fmt.Println(fl.MapsTo)
			cmdArgs = append(cmdArgs, "--"+fl.MapsTo) // append flag to arguments
		}
	}

	// run the command
	cmd := exec.Command(executablePath, cmdArgs...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
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

// Add the cobra commands in Host CLI
func LoadCommands(cmd *cobra.Command) {
	dateCommand := createCommandNew(&CommandConfig{
		Name:             "whatsup",
		MapsTo:           ArgsConfig{Name: "date", Subcommand: "-u", Args: []string{}},
		Flags:            []FlagConfig{},
		ShortDescription: "Tells date time",
		Examples:         "$ host whatsup",
	})

	yarnCommand := createCommandNew(&CommandConfig{
		Name:   "install",
		MapsTo: ArgsConfig{Name: "yarn", Subcommand: "add", Args: []string{"package-name"}},
		Flags: []FlagConfig{
			{
				Name:         "debug",
				MapsTo:       "verbose",
				Description:  "Debug mode",
				Alias:        "d",
				Type:         "bool",
				DefaultValue: "false",
			},
		},
		ShortDescription: "Install a package",
		Examples:         "$ host install",
	})

	addCommand := createCommandNew(&CommandConfig{
		Name:             "plus",
		MapsTo:           ArgsConfig{Name: "./plugins/calc", Subcommand: "add", Args: []string{"num1", "num2"}},
		ShortDescription: "Adds Two Nums",
		Examples:         "$ host plus 1 2",
		Flags: []FlagConfig{
			{
				Name:         "fl",
				Alias:        "f",
				DefaultValue: "false",
				Type:         "bool",
				Description:  "Add floating numbers",
				MapsTo:       "float",
			},
		},
	})

	gitcloneCommand := createCommandNew(&CommandConfig{
		Name:             "copy",
		MapsTo:           ArgsConfig{Name: "git", Subcommand: "clone", Args: []string{"repository_url"}},
		ShortDescription: "Clone a git repository",
		Examples:         "$ host copy https://github.com/ankithans/codeX",
		Flags: []FlagConfig{
			{
				Name:         "debug",
				MapsTo:       "verbose",
				Description:  "Debug mode",
				Alias:        "d",
				Type:         "bool",
				DefaultValue: "false",
			},
		},
	})

	cmd.AddCommand(dateCommand)
	cmd.AddCommand(yarnCommand)
	cmd.AddCommand(addCommand)
	cmd.AddCommand(gitcloneCommand)
}
