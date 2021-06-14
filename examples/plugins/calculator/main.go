package calculator

import (
	"strconv"

	"github.com/aerogear/charmil/pkg/factory"
	"github.com/aerogear/charmil/pkg/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	Logger func() (logging.Logger, error)
}

func RootCommand(f *factory.Factory) *cobra.Command {
	opts := &Options{
		Logger: f.Logger,
	}
	cmd := &cobra.Command{
		Use:          "calc",
		Short:        "Calculator",
		Example:      "$ host calc add 2 3",
		SilenceUsage: true,
	}

	cmd.AddCommand(AddCmd(opts))
	return cmd
}

func AddCmd(opts *Options) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Add numbers",
		Long:  `Adding the numbers`,
		Run: func(cmd *cobra.Command, args []string) {
			fstatus, _ := cmd.Flags().GetBool("float")
			if fstatus {
				addFloat(args, opts)
			} else {
				addInt(args, opts)
			}
		},
	}

	addCmd.Flags().BoolP("float", "f", false, "Add Floating Numbers")
	return addCmd
}

func addInt(args []string, opts *Options) {
	logger, _ := opts.Logger()
	var sum int
	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)
		if err != nil {
			logger.Error(err)
		}
		sum = sum + itemp
	}
	logger.Output("Addition of numbers", args, "is", sum)
}

func addFloat(args []string, opts *Options) {
	logger, _ := opts.Logger()
	var sum float64
	for _, fval := range args {
		ftemp, err := strconv.ParseFloat(fval, 64)
		if err != nil {
			logger.Error(err)
		}
		sum = sum + ftemp
	}
	logger.Output("Addition of numbers", args, "is", sum)
}
