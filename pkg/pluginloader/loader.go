package pluginloader

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"gopkg.in/yaml.v2"
)

func getFileList(p string) ([]string, error) {
	searchDir := p

	fileList := make([]string, 0)
	e := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, f.Name())
		return err
	})

	if e != nil {
		panic(e)
	}

	return fileList, nil
}

func AddCommandsNew(cmd *cobra.Command) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	flist, err := getFileList(path.Join(cwd, "./plugins"))
	if err != nil {
		return err
	}
	fmt.Print(flist)

	for i, files := range flist {
		if i == 0 {
		} else {
			b, err := ioutil.ReadFile(path.Join(cwd, "./plugins/"+files))
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
	}
	return nil
}

func AddCommands(cmd *cobra.Command) error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	b, err := ioutil.ReadFile(path.Join(cwd, "./plugins/git.yaml"))
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

			args = append([]string{cmdCfg.MapsTo.Subcommand}, args...)
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
