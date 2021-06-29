package rules

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/aerogear/charmil/validator"
	"github.com/spf13/cobra"
)

// define errors for the Length Rule
var (
	ErrLengthMin           = errors.New("less than min")
	ErrLengthMax           = errors.New("less than max")
	ErrLengthInvalid       = errors.New("invalid length")
	ErrLengthNegative      = errors.New("negative value")
	ErrLengthZeroValue     = errors.New("zero value")
	ErrLengthFieldNotExist = errors.New("field doesn't exists")
)

var LengthRule = "LENGTH_RULE"

// Length is a struct that provides a map
// with key as attribute for which length is controlled
// and value limit as Limit struct
type Length struct {
	Enable bool
	Limits map[string]Limit
}

// Limit defines min, max length of string
type Limit struct {
	Min, Max int
}

func validateLength(cmd *cobra.Command, l Length, verbose bool) []validator.ValidationError {
	var errors []validator.ValidationError

	for fieldName, limits := range l.Limits {
		// reflects the fieldName in cobra.Command struct
		reflectValue := reflect.ValueOf(cmd).Elem().FieldByName(fieldName)

		// if the defined fieldName doesn't exist in cobra.Command
		if reflectValue.String() == "<invalid Value>" {
			errors = append(errors, validator.ValidationError{Name: fmt.Sprintf("%s Field doesn't exist in cobra.Command", fieldName), Err: ErrLengthFieldNotExist, Rule: LengthRule, Cmd: cmd})
			continue
		}

		// validate fieldName
		err := validateField(cmd, limits, reflectValue.String(), cmd.CommandPath(), fieldName, verbose)
		if err.Err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}

// validateField compares the defined limit
// with the length of the attribute/value
func validateField(cmd *cobra.Command, limit Limit, value string, path string, fieldName string, verbose bool) validator.ValidationError {
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
		return validator.ValidationError{Name: fmt.Sprintf("%s length should be at least %d in %s cmd", fieldName, limit.Min, path), Err: ErrLengthMin, Rule: LengthRule, Cmd: cmd}
	}
	if length > limit.Max {
		return validator.ValidationError{Name: fmt.Sprintf("%s length should be less than %d in %s cmd", fieldName, limit.Max, path), Err: ErrLengthMax, Rule: LengthRule, Cmd: cmd}
	}

	return validator.ValidationError{}
}

// isLimitSet checks if the limit set
// by the user is valid or not
func isLimitSet(cmd *cobra.Command, limit Limit) (bool, validator.ValidationError) {
	if limit.Max < 0 || limit.Min < 0 {
		return true, validator.ValidationError{Name: "max and min must be greater than 0", Err: ErrLengthNegative, Rule: LengthRule, Cmd: cmd}
	}
	if limit.Max == 0 && limit.Min == 0 {
		return false, validator.ValidationError{Name: "limit not set", Err: ErrLengthZeroValue, Rule: LengthRule, Cmd: cmd}
	}
	if limit.Max < limit.Min {
		return true, validator.ValidationError{Name: "max limit must be greater than min limit", Err: ErrLengthMin, Rule: LengthRule, Cmd: cmd}
	}

	return true, validator.ValidationError{}
}
