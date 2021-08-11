package examplebuilder

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// NewCmdExample adds example with description to the cobra.Command
func NewCmdExample(cmd *cobra.Command, description string, flags []string) {
	cmd.Example += fmt.Sprintf("\n\n# %s \n%s %s", description, cmd.CommandPath(), strings.Join(flags, ", "))
}
