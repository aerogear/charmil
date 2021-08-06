package flagutil

import "github.com/spf13/cobra"

var (
	// ValidOutputFormats stores a default set of possible valid output formats
	ValidOutputFormats = []string{"json", "yml", "yaml", "toml"}

	// CredentialsOutputFormats stores a default set of possible valid credential output formats
	CredentialsOutputFormats = []string{"env", "json", "properties"}
)

// EnableStaticFlagCompletion enables autocompletion for flags with predefined valid values
func EnableStaticFlagCompletion(cmd *cobra.Command, flagName string, validValues []string) {
	_ = cmd.RegisterFlagCompletionFunc(flagName, func(cmd *cobra.Command, _ []string, _ string) ([]string, cobra.ShellCompDirective) {
		return validValues, cobra.ShellCompDirectiveNoSpace
	})
}
