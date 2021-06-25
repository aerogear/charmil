package validator

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

type LengthRule struct {
	Limits map[string]Limit
}

type Limit struct {
	Min, Max int
}

func (l *LengthRule) Validate(cmd *cobra.Command) []error {

	var errors []error
	err := validateCobraCommand(cmd, l)
	errors = append(errors, err...)

	for _, child := range cmd.Commands() {
		err := validateCobraCommand(child, l)
		errors = append(errors, err...)

	}

	return errors
}

func validateCobraCommand(cmd *cobra.Command, l *LengthRule) []error {
	var errors []error

	for fieldName, limits := range l.Limits {
		reflectValue := reflect.ValueOf(cmd).Elem().FieldByName(fieldName).String()

		err := validateField(limits, reflectValue, cmd.CommandPath())
		if err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

func validateField(limit Limit, value string, path string) error {
	length := len(value)

	_, err := isLimitSet(limit)
	if err != nil {
		return err
	}

	if length < limit.Min {
		return fmt.Errorf("%s in %s: length should be atleast %d", value, path, limit.Min)
	}
	if length > limit.Max {
		return fmt.Errorf("%s in %s: length should be less than %d", value, path, limit.Max)
	}

	return nil
}

func isLimitSet(limit Limit) (bool, error) {
	if limit.Max < 0 || limit.Min < 0 {
		return true, fmt.Errorf("max and min must be greater than 0")
	}
	if limit.Max == 0 && limit.Min == 0 {
		return false, fmt.Errorf("limit not set")
	}
	if limit.Max < limit.Min {
		return true, fmt.Errorf("max limit must be greater than min limit")
	}

	return true, nil
}
