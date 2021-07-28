package rules

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aerogear/charmil/validator"
	"github.com/spf13/cobra"
)

// define errors for the UseMatches Rule
var (
	ErrExampleMistmatch = errors.New("example mismatch")
)

var ExampleMatchesRule = "EXAMPLE_MATCHES_RULE"

// ExampleMatches defines Regexp to be compared
type ExampleMatches struct {
	RuleOptions validator.RuleOptions
}

// Validate is a method of type Rule Interface
// which returns validation errors
// checks if the example is updated as per command
func (e *ExampleMatches) Validate(cmd *cobra.Command) []validator.ValidationError {
	var errors []validator.ValidationError

	// if command needs to be ignored
	if val, ok := e.RuleOptions.SkipCommands[cmd.CommandPath()]; ok {
		if val {
			return errors
		}
	}

	cmdPath := cmd.CommandPath()
	cmdExample := cmd.Example

	if e.RuleOptions.Verbose {
		fmt.Fprintf(os.Stderr, "%s Command %s -> Example:%s, Path:%s\n", ExampleMatchesRule, cmd.CommandPath(), cmdExample, cmdPath)
	}

	// check if cmdPath is not the substring of cmdExample
	if cmdPath != "" && cmdExample != "" && !strings.Contains(cmdExample, cmdPath) {
		errors = append(errors, validator.ValidationError{Name: "provided example doesn't contain command path", Err: ErrExampleMistmatch, Rule: ExampleMatchesRule, Cmd: cmd})
	}

	return errors
}
