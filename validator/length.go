package validator

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

type Length struct {
	Limits map[string]Limit
}

type Limit struct {
	Min, Max int
}

func (l *Length) Validate(cmd *cobra.Command, verbose bool) []Error {

	var errors []Error
	info := stats{num: 0, num_failed: 0, errors: errors}

	err := validateLength(cmd, l, verbose)
	info.num++
	info.num_failed += len(err)
	errors = append(errors, err...)

	for _, child := range cmd.Commands() {
		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}

		err := validateLength(child, l, verbose)
		info.num++
		info.num_failed += len(err)
		errors = append(errors, err...)

	}

	if verbose {
		fmt.Printf("commands checked: %d\nchecks failed: %d\n", info.num, info.num_failed)
	}

	return errors
}

func validateLength(cmd *cobra.Command, l *Length, verbose bool) []Error {
	var errors []Error

	for fieldName, limits := range l.Limits {
		reflectValue := reflect.ValueOf(cmd).Elem().FieldByName(fieldName).String()

		err := validateField(limits, reflectValue, cmd.CommandPath(), verbose)
		if err.Err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func validateField(limit Limit, value string, path string, verbose bool) Error {
	length := len(value)

	_, err := isLimitSet(limit)
	if err.Err != nil {
		return err
	}

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
