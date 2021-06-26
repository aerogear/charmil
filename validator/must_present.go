package validator

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

// MustPresent is a struct that provides
// Fields defined for MustPresent validation
type MustPresent struct {
	Fields []string
}

// Validate implements the Rule interface
func (p *MustPresent) Validate(cmd *cobra.Command, verbose bool) []Error {
	var errors []Error
	info := stats{num: 0, num_failed: 0, errors: errors}

	return p.ValidateHelper(cmd, verbose, info)
}

// ValidateHelper returns errors of type []Error
// if the required field is not present
func (p *MustPresent) ValidateHelper(cmd *cobra.Command, verbose bool, info stats) []Error {

	// validate the root command
	err := validateMustPresent(cmd, p, verbose)
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
		if err := p.ValidateHelper(child, verbose, info); err != nil {
			return err
		}
	}

	// prints additional info in debug mode
	if verbose {
		fmt.Printf("commands checked: %d\nchecks failed: %d\n", info.num, info.num_failed)
	}

	return info.errors
}

// validateMustPresent traverse over all the fields
func validateMustPresent(cmd *cobra.Command, p *MustPresent, verbose bool) []Error {
	var errors []Error

	for _, field := range p.Fields {
		// reflects the field in cobra.Command struct
		reflectValue := reflect.ValueOf(cmd).Elem().FieldByName(field)

		// if the defined field doesn't exist in cobra.Command
		if reflectValue.String() == "<invalid Value>" {
			errors = append(errors, Error{Name: fmt.Sprintf("%s Field doesn't exist in cobra.Command", field), Err: ErrFieldNotExist, Rule: LengthRule})
			continue
		}

		// validate field and append errors
		errors = append(errors, validateByType(&reflectValue, field, cmd.CommandPath(), verbose)...)
	}
	return errors
}

// ValidateByType handles all types of attributes
// provided by cobra.Command struct
func validateByType(reflectValue *reflect.Value, field string, path string, verbose bool) []Error {
	var errors []Error

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
		errors = append(errors, Error{Name: fmt.Sprintf("%s must be present in %s", field, path), Err: ErrAbsent, Rule: MustPresentRule})
	}

	return errors
}
