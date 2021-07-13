package rules

import (
	"errors"
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
// compares the regexp with Use attribute
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

	// check if cmdPath is not the substring of cmdExample
	if cmdPath != "" && cmdExample != "" && !strings.Contains(cmdExample, cmdPath) {
		errors = append(errors, validator.ValidationError{Name: "provided example doesn't match with command", Err: ErrExampleMistmatch, Rule: ExampleMatchesRule, Cmd: cmd})
	}

	return errors
}
