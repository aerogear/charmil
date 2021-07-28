package rules

import (
	"errors"
	"fmt"
	"os"

	"github.com/aerogear/charmil/validator"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// define errors for the UseMatches Rule
var (
	ErrFlagFullStop = errors.New("flag full stop")
)

var PunctuationRule = "PUNCTUATION_RULE"

// UseMatches defines Regexp to be compared
type Punctuation struct {
	RuleOptions validator.RuleOptions
}

// Validate is a method of type Rule Interface
// which returns validation errors
func (p *Punctuation) Validate(cmd *cobra.Command) []validator.ValidationError {
	var errors []validator.ValidationError

	// if command needs to be ignored
	if val, ok := p.RuleOptions.SkipCommands[cmd.CommandPath()]; ok {
		if val {
			return errors
		}
	}

	flags := cmd.Flags()
	flags.VisitAll(func(f *pflag.Flag) {
		usage := f.Usage

		// verbose logging
		if p.RuleOptions.Verbose {
			fmt.Fprintf(os.Stderr, "%s Command %s -> Flag - %s: usage - %s\n", PunctuationRule, cmd.CommandPath(), f.Name, usage)
		}

		// check for full-stops behind flags usage strings
		if usage[len(usage)-1] == 46 {
			errors = append(errors, validator.ValidationError{Name: "full stop in the flag usage", Err: ErrFlagFullStop, Rule: PunctuationRule, Cmd: cmd})
		}
	})

	return errors
}
