package commands

import (
	"errors"
	"github.com/spf13/cobra"
)

var (
	ErrNotInDevEnviroment = errors.New("cannot run developer command in this build")
)

func CharmilCommands() (*cobra.Command, error) {

	if !inDev {
		return nil, ErrNotInDevEnviroment
	}

	cmd := &cobra.Command{
		Use:     "charmil",
		Short:   "built in charmil commands",
		SilenceUsage: false,
	}

	validateCommand, err := ValidateCommand()
	if err != nil {
		return nil, err
	}

	cmd.AddCommand(validateCommand)

	return cmd, nil
}
