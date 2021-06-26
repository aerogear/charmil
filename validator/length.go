package validator

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

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
func (l *Length) Validate(cmd *cobra.Command, verbose bool) []Error {
	var errors []Error
	info := stats{num: 0, num_failed: 0, errors: errors}

	return l.ValidateHelper(cmd, verbose, info)
}

// ValidateHelper returns errors of type []Error
// if length is not validated
func (l *Length) ValidateHelper(cmd *cobra.Command, verbose bool, info stats) []Error {

	// validate the root command
	err := validateLength(cmd, l, verbose)
	// record stats
	info.num++
	info.num_failed += len(err)
	info.errors = append(info.errors, err...)

	// traverse descendents of cmd
	for _, child := range cmd.Commands() {

		// base case
		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}

		// recursive call for ValidateHelper
		if err := l.ValidateHelper(child, verbose, info); err != nil {
			return err
		}
	}

	// prints additional info in debug mode
	if verbose {
		fmt.Printf("commands checked: %d\nchecks failed: %d\n", info.num, info.num_failed)
	}

	return info.errors
}

// ValidateLength traverse over all the fields
// provided and compare with the limit
func validateLength(cmd *cobra.Command, l *Length, verbose bool) []Error {
	var errors []Error

	for fieldName, limits := range l.Limits {
		// reflects the fieldName in cobra.Command struct
		reflectValue := reflect.ValueOf(cmd).Elem().FieldByName(fieldName)

		// if the defined fieldName doesn't exist in cobra.Command
		if reflectValue.String() == "<invalid Value>" {
			errors = append(errors, Error{Name: fmt.Sprintf("%s Field doesn't exist in cobra.Command", fieldName), Err: ErrFieldNotExist, Rule: LengthRule})
			continue
		}

		// validate fieldName
		err := validateField(limits, reflectValue.String(), cmd.CommandPath(), verbose)
		if err.Err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

// validateField compares the defined limit
// with the length of the attribute/value
func validateField(limit Limit, value string, path string, verbose bool) Error {
	length := len(value)

	// check if valid limit is set
	_, err := isLimitSet(limit)
	if err.Err != nil {
		return err
	}

	// prints additional info in debug mode
	if verbose {
		fmt.Printf("%s Command %s -> %s: %v\n", LengthRule, path, value, limit)
	}

	if length < limit.Min {
		return Error{Name: fmt.Sprintf("length should be at least %d", limit.Min), Err: ErrMin, Rule: LengthRule}
	}
	if length > limit.Max {
		return Error{Name: fmt.Sprintf("length should be less than %d", limit.Max), Err: ErrMax, Rule: LengthRule}
	}

	return Error{}
}

// isLimitSet checks if the limit set
// by the user is valid or not
func isLimitSet(limit Limit) (bool, Error) {
	if limit.Max < 0 || limit.Min < 0 {
		return true, Error{Name: "max and min must be greater than 0", Err: ErrNeg, Rule: LengthRule}
	}
	if limit.Max == 0 && limit.Min == 0 {
		return false, Error{Name: "limit not set", Err: ErrZeroValue, Rule: LengthRule}
	}
	if limit.Max < limit.Min {
		return true, Error{Name: "max limit must be greater than min limit", Err: ErrMin, Rule: LengthRule}
	}

	return true, Error{}
}
