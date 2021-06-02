package pluginloader

import (
	"bytes"
	"io"
	"os"
	"os/exec"

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

	// TODO: add Flags

	return cmd
}

// Add the cobra command in Host CLI
func LoadCommands(cmd *cobra.Command) {
	dateCommand := createCommand(&CommandConfig{
		Name:             "whatsup",
		MapsTo:           ArgsConfig{Name: "date", Subcommand: "-u", Args: []string{}},
		Flags:            []FlagConfig{},
		ShortDescription: "Tells date timw",
		Examples:         "$ host whatsup",
	})

	yarnCommand := createCommand(&CommandConfig{
		Name:             "install",
		MapsTo:           ArgsConfig{Name: "yarn", Subcommand: "add", Args: []string{"package-name"}},
		Flags:            []FlagConfig{},
		ShortDescription: "Install a package",
		Examples:         "$ host install",
	})

	cmd.AddCommand(dateCommand)
	cmd.AddCommand(yarnCommand)
}
