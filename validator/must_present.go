package validator

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

type MustPresent struct {
	Fields []string
}

func (p *MustPresent) Validate(cmd *cobra.Command, verbose bool) []Error {
	var errors []Error
	info := stats{num: 0, num_failed: 0, errors: errors}

	return p.ValidateHelper(cmd, verbose, info)
}

func (p *MustPresent) ValidateHelper(cmd *cobra.Command, verbose bool, info stats) []Error {

	err := validateMustPresent(cmd, p, verbose)
	info.num++
	info.num_failed += len(err)
	info.errors = append(info.errors, err...)

	for _, child := range cmd.Commands() {

		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}

		if err := p.ValidateHelper(child, verbose, info); err != nil {
			return err
		}
	}

	if verbose {
		fmt.Printf("commands checked: %d\nchecks failed: %d\n", info.num, info.num_failed)
	}

	return info.errors
}

func validateMustPresent(cmd *cobra.Command, p *MustPresent, verbose bool) []Error {
	var errors []Error

	for _, field := range p.Fields {
		reflectValue := reflect.ValueOf(cmd).Elem().FieldByName(field)
		errors = append(errors, validateByType(&reflectValue, field, cmd.CommandPath(), verbose)...)
	}
	return errors
}

func validateByType(reflectValue *reflect.Value, field string, path string, verbose bool) []Error {
	var errors []Error

	if verbose {
		fmt.Printf("%s Command %s -> %s: %v\n", MustPresentRule, path, reflectValue.String(), field)
	}

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
