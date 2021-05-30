package pluginloader

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// wrapping the yarn command and creating a cobra command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "install a package",
	Long:  ``,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Install command called")

		args = append([]string{"add"}, args...)
		fmt.Println(args)

		c := exec.Command("yarn", args...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		var buf bytes.Buffer
		c.Stderr = io.MultiWriter(os.Stderr, &buf)

		return c.Run()
	},
}
