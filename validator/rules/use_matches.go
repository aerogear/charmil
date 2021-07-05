package rules

import (
	"errors"
	"regexp"

	"github.com/aerogear/charmil/validator"
	"github.com/spf13/cobra"
)

// define errors for the UseMatches Rule
var (
	ErrInvalidRegexp = errors.New("invalid regexp")
	ErrRegexMismatch = errors.New("regexp mismatch")
)

var UseMatchesRule = "USE_MATCHES_RULE"

// UseMatches defines Regexp to be compared
type UseMatches struct {
	Regexp string `json:"Regexp"`
}

// Validate is a method of type Rule Interface
// which returns validation errors
func (u *UseMatches) Validate(cmd *cobra.Command) []validator.ValidationError {
	return u.validateUse(cmd)
}

// validateUse compares the regexp with Use attribute
func (u *UseMatches) validateUse(cmd *cobra.Command) []validator.ValidationError {
	var errors []validator.ValidationError

	r, err := regexp.Compile(u.Regexp)
	if err != nil {
		errors = append(errors, validator.ValidationError{Name: "given regexp is invalid", Err: ErrInvalidRegexp, Rule: UseMatchesRule, Cmd: cmd})
		return errors
	}
	if !r.MatchString(cmd.Use) {
		errors = append(errors, validator.ValidationError{Name: "Use doesn't match with given regexp", Err: ErrRegexMismatch, Rule: UseMatchesRule, Cmd: cmd})
	}

	return errors
}
