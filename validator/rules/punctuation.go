package rules

import (
	"github.com/aerogear/charmil/validator"
	"github.com/spf13/cobra"
)

// define errors for the UseMatches Rule
var ()

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

	// punctuation logic

	return errors
}
