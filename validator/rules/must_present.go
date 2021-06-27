package rules

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

var MustPresentRule = "MUST_PRESENT_RULE"

// MustPresent is a struct that provides
// Fields defined for MustPresent validation
type MustPresent struct {
	Fields []string
}

// Validate implements the Rule interface
func (p *MustPresent) ValidateMustPresent(cmd *cobra.Command, verbose bool) []ValidationError {
	var errors []ValidationError
	info := StatusLog{totalTested: 0, totalErrors: 0, errors: errors}

	return Traverse(
		cmd,
		verbose,
		info,
		p,
		func(cmd *cobra.Command, verbose bool) []ValidationError {
			return p.validateHelper(cmd, verbose)
		},
	)
}

func (p *MustPresent) validateHelper(cmd *cobra.Command, verbose bool) []ValidationError {
	var errors []ValidationError

	for _, field := range p.Fields {
		// reflects the field in cobra.Command struct
		reflectValue := reflect.ValueOf(cmd).Elem().FieldByName(field)

		// if the defined field doesn't exist in cobra.Command
		if reflectValue.String() == "<invalid Value>" {
			errors = append(errors, ValidationError{Name: fmt.Sprintf("%s Field doesn't exist in cobra.Command", field), Err: ErrFieldNotExist, Rule: LengthRule, cmd: cmd})
			continue
		}

		// validate field and append errors
		errors = append(errors, validateByType(cmd, &reflectValue, field, cmd.CommandPath(), verbose)...)
	}
	return errors
}

// ValidateByType handles all types of attributes
// provided by cobra.Command struct
func validateByType(cmd *cobra.Command, reflectValue *reflect.Value, field string, path string, verbose bool) []ValidationError {
	var errors []ValidationError

	// prints additional info in debug mode
	if verbose {
		fmt.Printf("%s Command %s -> %s: %v\n", MustPresentRule, path, reflectValue.String(), field)
	}

	// handle types
	if reflectValue.String() == "" ||
		(reflectValue.Kind().String() == "func" && reflectValue.IsNil()) ||
		(reflectValue.Kind().String() == "bool" && !reflectValue.Bool()) ||
		(reflectValue.Kind().String() == "int" && reflectValue.Int() == 0) ||
		(reflectValue.Kind().String() == "slice" && reflectValue.Len() == 0) ||
		(reflectValue.Kind().String() == "map" && reflectValue.Len() == 0) {
		errors = append(errors, ValidationError{Name: fmt.Sprintf("%s must be present in %s", field, path), Err: ErrAbsent, Rule: MustPresentRule, cmd: cmd})
	}

	return errors
}
