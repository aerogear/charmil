package validator

import (
	"fmt"
	"reflect"

	"github.com/spf13/cobra"
)

type MustPresentRule struct {
	Fields []string
}

func (p *MustPresentRule) Validate(cmd *cobra.Command) []error {

	var errors []error
	err := validateMustPresent(cmd, p)
	errors = append(errors, err...)

	for _, child := range cmd.Commands() {
		err := validateMustPresent(child, p)
		errors = append(errors, err...)
	}

	return errors
}

func validateMustPresent(cmd *cobra.Command, p *MustPresentRule) []error {
	var errors []error

	for _, field := range p.Fields {
		reflectValue := reflect.ValueOf(cmd).Elem().FieldByName(field)
		errors = append(errors, validateByType(&reflectValue, field, cmd.CommandPath())...)
	}
	return errors
}

func validateByType(reflectValue *reflect.Value, field string, path string) []error {
	var errors []error

	if reflectValue.String() == "" ||
		(reflectValue.Kind().String() == "func" && reflectValue.IsNil()) ||
		(reflectValue.Kind().String() == "bool" && !reflectValue.Bool()) ||
		(reflectValue.Kind().String() == "int" && reflectValue.Int() == 0) ||
		(reflectValue.Kind().String() == "slice" && reflectValue.Len() == 0) ||
		(reflectValue.Kind().String() == "map" && reflectValue.Len() == 0) {
		errors = append(errors, fmt.Errorf("%s must be present in %s", field, path))
	}

	return errors
}
