package core

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

type PluginConfig struct {
	Commands []CommandConfig `yaml:"commands"`
}

// CommandConfig
type CommandConfig struct {
	Name             string        `yaml:"name"`
	MapsTo           ArgsMapConfig `yaml:"mapsTo"`
	Flags            []FlagConfig  `yaml:"flags"`
	ShortDescription string        `yaml:"shortDescription"`
	Examples         string        `yaml:"usage"`
}

type ArgsMapConfig struct {
	Name       string
	Subcommand string
	Args       []string
}

// FlagConfig
type FlagConfig struct {
	Type         string `yaml:"type"`
	DefaultValue string `yaml:"defaultValue"`
	Name         string `yaml:"name"`
	MapsTo       string `yaml:"mapsTo"`
	Description  string `yaml:"description"`
	Alias        string `yaml:"alias"`
}

func LoadCommands(cmd *cobra.Command) error {
	pluginFilenames, err := ioutil.ReadDir("./plugins")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range pluginFilenames {
		err = newCommands(cmd, f.Name())
		if err != nil {
			log.Fatal(err)
		}
	}
	return nil
}

func newCommands(cmd *cobra.Command, pluginFilename string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	b, err := ioutil.ReadFile(path.Join(cwd, "./plugins/"+pluginFilename))
	if err != nil {
		return err
	}

	var cliPlugin *PluginConfig
	err = yaml.Unmarshal(b, &cliPlugin)
	if err != nil {
		return err
	}

	if &cliPlugin.Commands != nil && len(cliPlugin.Commands) > 0 {
		for _, cfg := range cliPlugin.Commands {
			cmd.AddCommand(addCommand(&cfg))
		}
	}
	return nil
}

func addCommand(cmdCfg *CommandConfig) *cobra.Command {
	cmd := &cobra.Command{
		Use:           cmdCfg.Name,
		Short:         cmdCfg.ShortDescription,
		SilenceErrors: true,
		Example:       cmdCfg.Examples,
		Args:          cobra.ExactArgs(len(cmdCfg.MapsTo.Args)),
		RunE: func(cmd *cobra.Command, args []string) error {
			if cmdCfg.MapsTo.Subcommand != "" {
				args = append([]string{cmdCfg.MapsTo.Subcommand}, args...)
			}
			c := exec.Command(cmdCfg.MapsTo.Name, args...)
			c.Stdin = os.Stdin
			c.Stdout = os.Stdout
			var buf bytes.Buffer
			c.Stderr = io.MultiWriter(os.Stderr, &buf)

			return c.Run()
		},
	}

	if cmdCfg.Flags != nil && len(cmdCfg.Flags) > 0 {
		for _, f := range cmdCfg.Flags {
			fs := cmd.Flags()
			addFlag(&f, fs)
		}
	}

	return cmd
}

func addFlag(flagCfg *FlagConfig, fs *pflag.FlagSet) {
	switch flagCfg.Type {
	case "string":
		fs.StringP(flagCfg.Name, flagCfg.Alias, flagCfg.DefaultValue, flagCfg.Description)
	case "bool":
		v, _ := strconv.ParseBool(flagCfg.DefaultValue)
		fs.BoolP(flagCfg.Name, flagCfg.Alias, v, flagCfg.Description)
	case "int":
		v, _ := strconv.Atoi(flagCfg.DefaultValue)
		fs.IntP(flagCfg.Name, flagCfg.Alias, v, flagCfg.Description)
	}
}
