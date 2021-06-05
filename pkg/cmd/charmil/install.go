package charmil

import (
	"bytes"
	"io"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

// InstallCmd represents the install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Install plugin(s)",
	Long:  ``,
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// Temporary code. Will be later replaced with logic which downloads, extracts, installs and validates binary from manifest file

		pluginName := args[0]
		c := exec.Command("go", "build", "-buildmode=plugin", "-o=./plugins/shared_objects", "./plugins/"+pluginName)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		var buf bytes.Buffer
		c.Stderr = io.MultiWriter(os.Stderr, &buf)

		return c.Run()
	},
}
