package calculator

import (
	"fmt"
	"strconv"

	"github.com/aerogear/charmil/pkg/factory"
	"github.com/aerogear/charmil/pkg/localize"
	"github.com/aerogear/charmil/pkg/logging"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
)

type Options struct {
	Logger   func() (logging.Logger, error)
	Localize localize.Localizer
}

func RootCommand() *cobra.Command {

	loc, err := localize.InitLocalizer(localize.Config{Language: language.English, Path: "examples/plugins/calculator/locals/en/en.yaml", Format: "yaml"})

	if err != nil {
		fmt.Println("Error", err)
	}

	newFactory := factory.Default(loc)
	opts := &Options{
		Logger:   newFactory.Logger,
		Localize: newFactory.Localizer,
	}
	cmd := &cobra.Command{
		Use:          opts.Localize.LocalizeByID("calc.cmd.use"),
		Short:        opts.Localize.LocalizeByID("calc.cmd.short"),
		Example:      opts.Localize.LocalizeByID("calc.cmd.example"),
		SilenceUsage: true,
	}

	cmd.AddCommand(AddCmd(opts))
	return cmd
}

func AddCmd(opts *Options) *cobra.Command {
	addCmd := &cobra.Command{
		Use:   opts.Localize.LocalizeByID("add.cmd.use"),
		Short: opts.Localize.LocalizeByID("add.cmd.short"),
		Long:  opts.Localize.LocalizeByID("add.cmd.long"),
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
