package rules

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

var LengthRule = "LENGTH_RULE"

// Length is a struct that provides a map
// with key as attribute for which length is controlled
// and value limit as Limit struct
type Length struct {
	Limits map[string]Limit
}

// Limit defines min, max length of string
type Limit struct {
	Min, Max int
}

// Validate implements the Rule interface
func (l *Length) ValidateLength(cmd *cobra.Command, verbose bool) []ValidationError {
	var errors []ValidationError
	info := StatusLog{totalTested: 0, totalErrors: 0, errors: errors}

	return Traverse(
		cmd,
		verbose,
		info,
		l,
		func(cmd *cobra.Command, verbose bool) []ValidationError {
			return l.validateHelper(cmd, verbose)
		},
	)
}

func (l *Length) validateHelper(cmd *cobra.Command, verbose bool) []ValidationError {
	var errors []ValidationError

	for fieldName, limits := range l.Limits {
		// reflects the fieldName in cobra.Command struct
		reflectValue := reflect.ValueOf(cmd).Elem().FieldByName(fieldName)

		// if the defined fieldName doesn't exist in cobra.Command
		if reflectValue.String() == "<invalid Value>" {
			errors = append(errors, ValidationError{Name: fmt.Sprintf("%s Field doesn't exist in cobra.Command", fieldName), Err: ErrFieldNotExist, Rule: LengthRule, cmd: cmd})
			continue
		}

		// validate fieldName
		err := validateField(cmd, limits, reflectValue.String(), cmd.CommandPath(), verbose)
		if err.Err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

// validateField compares the defined limit
// with the length of the attribute/value
func validateField(cmd *cobra.Command, limit Limit, value string, path string, verbose bool) ValidationError {
	length := len(value)

	// check if valid limit is set
	_, err := isLimitSet(cmd, limit)
	if err.Err != nil {
		return err
	}

	// prints additional info in debug mode
	if verbose {
		fmt.Printf("%s Command %s -> %s: %v\n", LengthRule, path, value, limit)
	}

	if length < limit.Min {
		return ValidationError{Name: fmt.Sprintf("length should be at least %d", limit.Min), Err: ErrMin, Rule: LengthRule, cmd: cmd}
	}
	if length > limit.Max {
		return ValidationError{Name: fmt.Sprintf("length should be less than %d", limit.Max), Err: ErrMax, Rule: LengthRule, cmd: cmd}
	}

	return ValidationError{}
}

// isLimitSet checks if the limit set
// by the user is valid or not
func isLimitSet(cmd *cobra.Command, limit Limit) (bool, ValidationError) {
	if limit.Max < 0 || limit.Min < 0 {
		return true, ValidationError{Name: "max and min must be greater than 0", Err: ErrNeg, Rule: LengthRule, cmd: cmd}
	}
	if limit.Max == 0 && limit.Min == 0 {
		return false, ValidationError{Name: "limit not set", Err: ErrZeroValue, Rule: LengthRule, cmd: cmd}
	}
	if limit.Max < limit.Min {
		return true, ValidationError{Name: "max limit must be greater than min limit", Err: ErrMin, Rule: LengthRule, cmd: cmd}
	}

	return true, ValidationError{}
}
