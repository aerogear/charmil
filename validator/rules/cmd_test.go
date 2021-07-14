package rules

import (
	"testing"

	"github.com/spf13/cobra"
)

func Test_ExecuteCommand(t *testing.T) {

	// Testing cobra commands with default recommended config
	// default config can also be overrided
	ruleCfg := ValidatorConfig{
		ValidatorOptions: ValidatorOptions{
			SkipChildren: map[string]bool{
				"root echo": true,
			},
		},
		ValidatorRules: ValidatorRules{
			Length: Length{
				Limits: map[string]Limit{
					"Use":     {Min: 1},
					"Example": {Min: 3},
					"Long":    {Min: 5},
				},
			},
			MustExist: MustExist{
				Fields: map[string]bool{"Run": true},
			},
			UseMatches: UseMatches{Regexp: `^[^-_+]+$`},
		},
	}

	validationErr := ExecuteRules(rootCmd, &ruleCfg)

	for _, errs := range validationErr {
		if errs.Err != nil {
			t.Errorf("%s: cmd %s: %s", errs.Rule, errs.Cmd.CommandPath(), errs.Name)
		}
	}

}

func emptyRun(*cobra.Command, []string) {}

func init() {
	echoCmd.AddCommand(timesCmd, echoSubCmd, printCmd)
	rootCmd.AddCommand(echoCmd, dummyCmd)
}

var rootCmd = &cobra.Command{
	Use:     "root",
	Short:   "Root short description",
	Long:    "Root long description",
	Example: "root",
	Run:     emptyRun,
}

var echoCmd = &cobra.Command{
	Use:     "echo [string to echo]",
	Aliases: []string{"say"},
	Short:   "Echo anything to the screen",
	Long:    "an utterly useless command for testing",
	Example: "Just run root echo",
	Run:     emptyRun,
}

var echoSubCmd = &cobra.Command{
	Use:     "echosub [string to print]",
	Short:   "second sub command for echo",
	Long:    "an absolutely utterly useless command for testing",
	Example: "root echo echosub",
	Run:     emptyRun,
}

var timesCmd = &cobra.Command{
	Use:        "times [# times] [string to echo]",
	SuggestFor: []string{"counts"},
	Short:      "Echo anything to the screen more times",
	Long:       `a slightly useless command for testing.`,
	Example:    "root echo times",
	Run:        emptyRun,
}

var printCmd = &cobra.Command{
	Use:     "print [string to print]",
	Short:   "Print anything to the screen",
	Long:    `an absolutely utterly useless command for testing.`,
	Example: "print",
	Run:     emptyRun,
}

var dummyCmd = &cobra.Command{
	Use:     "dummy [action]",
	Short:   "Performs a dummy action",
	Example: "dummy",
}
