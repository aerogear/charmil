package flagutil

import (
	"fmt"
)

// InvalidValueError returns an error when an invalid flag value is provided
func InvalidValueError(flag string, val interface{}, validOptions ...string) error {
	var chooseFromStr string
	if len(validOptions) > 0 {
		chooseFromStr = ", valid options are: "
		for i, option := range validOptions {
			chooseFromStr += fmt.Sprintf(`"%v"`, option)
			if (i + 1) < len(validOptions) {
				chooseFromStr += ", "
			}
		}
	}
	return fmt.Errorf(`invalid value "%v" for --%v%v`, val, flag, chooseFromStr)
}

// RequiredValueError returns an error when a required value is not provided
func RequiredValueError(arg string) error {
	return fmt.Errorf("%v requires a value", arg)
}

// CastValueError returns an error when a value cannot be casted to the given type
func CastValueError(v interface{}, t string) error {
	return fmt.Errorf(`could not cast %v, to type "%v"`, v, t)
}
