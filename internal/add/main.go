package add

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

func addInt(args []string) {
	var sum int
	for _, ival := range args {
		itemp, err := strconv.Atoi(ival)
		if err != nil {
			fmt.Println(err)
		}
		sum = sum + itemp
	}
	fmt.Printf("Addition of numbers %s is %d\n", args, sum)
}

// sample command for addition of 2 numbers
var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add two integers",
	Run: func(cmd *cobra.Command, args []string) {
		addInt(args)
	},
}
