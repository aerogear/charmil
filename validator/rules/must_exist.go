package rules

import (
	"errors"
	"fmt"
	"os"
	"reflect"

	"github.com/aerogear/charmil/validator"
	"github.com/spf13/cobra"
)

// define errors for the Must Exist Rule
var (
	ErrMustExistFieldNotExist = errors.New("field doesn't exists")
	ErrMustExistAbsent        = errors.New("field must be present")
)

var MustExistRule = "MUST_EXIST_RULE"

// MustExist is a struct that provides
// Fields defined for MustExist validation
type MustExist struct {
	RuleOptions validator.RuleOptions
	Fields      map[string]bool `json:"Fields"`
}

// Validate is a method of type Rule Interface
// which returns validation errors
func (m *MustExist) Validate(cmd *cobra.Command) []validator.ValidationError {
	var errors []validator.ValidationError

	// if command needs to be ignored
	if val, ok := m.RuleOptions.SkipCommands[cmd.CommandPath()]; ok {
		if val {
			return errors
		}
	}

	for field, isTrue := range m.Fields {
		// reflects the field in cobra.Command struct
		reflectValue := reflect.ValueOf(cmd).Elem().FieldByName(field)

		// if the defined field doesn't exist in cobra.Command
		if reflectValue.String() == "<invalid Value>" {
			errors = append(errors, validator.ValidationError{Name: fmt.Sprintf("%s Field doesn't exist in cobra.Command", field), Err: ErrMustExistFieldNotExist, Rule: LengthRule, Cmd: cmd})
			continue
		}

		// validate field and append errors
		if isTrue {
			errors = append(errors, validateByType(cmd, &reflectValue, field, cmd.CommandPath(), m.RuleOptions.Verbose)...)
		}
	}
	return errors
}

// ValidateByType handles all types of attributes
// provided by cobra.Command struct
func validateByType(cmd *cobra.Command, reflectValue *reflect.Value, field string, path string, verbose bool) []validator.ValidationError {
	var errors []validator.ValidationError

	// prints additional info in debug mode
	if verbose {
		fmt.Fprintf(os.Stderr, "%s Command %s -> %s: %v\n", MustExistRule, path, reflectValue.String(), field)
	}

	// handle types
	if reflectValue.String() == "" ||
		(reflectValue.Kind().String() == "func" && reflectValue.IsNil()) ||
		(reflectValue.Kind().String() == "bool" && !reflectValue.Bool()) ||
		(reflectValue.Kind().String() == "int" && reflectValue.Int() == 0) ||
		(reflectValue.Kind().String() == "slice" && reflectValue.Len() == 0) ||
		(reflectValue.Kind().String() == "map" && reflectValue.Len() == 0) {
		errors = append(errors, validator.ValidationError{Name: fmt.Sprintf("%s must be present", field), Err: ErrMustExistAbsent, Rule: MustExistRule, Cmd: cmd})
	}

	return errors
}
