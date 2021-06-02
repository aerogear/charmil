package pluginloader

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"strconv"

	"github.com/spf13/cobra"
)

// TODO:
// Unmarshal the YAML
// Extract the commands using plugin.go (command structure)

// Create a Cobra Command
func createCommand(cmdStruct *CommandConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:     cmdStruct.Name,
		Short:   cmdStruct.ShortDescription,
		Example: cmdStruct.Examples,
		Args:    cobra.ExactArgs(len(cmdStruct.MapsTo.Args)),
		RunE: func(cmd *cobra.Command, args []string) error {
			args = append([]string{cmdStruct.MapsTo.Subcommand}, args...)
			c := exec.Command(cmdStruct.MapsTo.Name, args...)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			var buf bytes.Buffer
			c.Stderr = io.MultiWriter(os.Stderr, &buf)
			return c.Run()
		},
	}

	// add Flags
	if cmdStruct.Flags != nil && len(cmdStruct.Flags) > 0 {
		for _, f := range cmdStruct.Flags {
			fs := cmd.Flags()
			switch f.Type {
			case "string":
				fs.StringP(f.Name, f.Alias, f.DefaultValue, f.Description)
			case "bool":
				v, _ := strconv.ParseBool(f.DefaultValue)
				fs.BoolP(f.Name, f.Alias, v, f.Description)
			case "int":
				v, _ := strconv.Atoi(f.DefaultValue)
				fs.IntP(f.Name, f.Alias, v, f.Description)
			}
		}
	}

	return cmd
}

// Add the cobra command in Host CLI
func LoadCommands(cmd *cobra.Command) {
	dateCommand := createCommand(&CommandConfig{
		Name:             "whatsup",
		MapsTo:           ArgsConfig{Name: "date", Subcommand: "-u", Args: []string{}},
		Flags:            []FlagConfig{},
		ShortDescription: "Tells date time",
		Examples:         "$ host whatsup",
	})

	yarnCommand := createCommand(&CommandConfig{
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

	cmd.AddCommand(dateCommand)
	cmd.AddCommand(yarnCommand)
}
