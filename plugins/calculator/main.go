package calculator

import (
	"fmt"
	"strconv"

	"github.com/aerogear/charmil/pkg/pluginloader"

	"github.com/spf13/cobra"
)

var AddCommand = pluginloader.CreateCommand(&pluginloader.CommandConfig{
	Name: "add",
	Args: []string{"num1", "num2"},
	Flags: []pluginloader.FlagConfig{{
		Name:         "float",
		Alias:        "f",
		DefaultValue: "false",
		Type:         "bool",
		Description:  "Add floating numbers",
	}},
	ShortDescription: "add numbers",
	Examples:         "$ host plus 1 2",
}, func(cmd1 *cobra.Command, args []string) {
	ExecuteAddition(cmd1, args)
})

func ExecuteAddition(cmd *cobra.Command, args []string) {
	fstatus, _ := cmd.Flags().GetBool("float")
	if fstatus {
		addFloat(args)
	} else {
		addInt(args)
	}
}

func addInt(args []string) {
	var sum int
	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)

		if err != nil {
			fmt.Println(err)
		}

		sum = sum + itemp
	}
	fmt.Printf("Addition of numbers %s is %d", args, sum)
}

func addFloat(args []string) {
	var sum float64
	for _, fval := range args {
		ftemp, err := strconv.ParseFloat(fval, 64)

		if err != nil {
			fmt.Println(err)
		}
		sum = sum + ftemp
	}
	fmt.Printf("Sum of floating numbers is %s is %f", args, sum)
}
