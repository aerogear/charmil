package validator

import (
	"fmt"

	"github.com/spf13/cobra"
)

type LengthRule struct {
	Use     Limit
	Short   Limit
	Long    Limit
	Example Limit
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

	cmdPath := cmd.CommandPath()

	use := cmd.Use
	useErr := validateField(l.Use, use, cmdPath)
	if useErr != nil {
		errors = append(errors, useErr)
	}

	short := cmd.Short
	shortErr := validateField(l.Short, short, cmdPath)
	if shortErr != nil {
		errors = append(errors, shortErr)
	}

	long := cmd.Long
	longErr := validateField(l.Long, long, cmdPath)
	if longErr != nil {
		errors = append(errors, longErr)
	}

	example := cmd.Example
	exampleErr := validateField(l.Example, example, cmdPath)
	if exampleErr != nil {
		errors = append(errors, exampleErr)
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
