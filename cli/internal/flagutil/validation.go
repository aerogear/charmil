package flagutil

import (
	"errors"
	"fmt"
	"regexp"
)

// ValidateName validates the proposed name of an instance
func ValidateName(val interface{}, validRegexpString string, minLength, maxLength int) error {

	name, ok := val.(string)
	if !ok {
		return CastValueError(val, "string")
	}

	if len(name) < minLength || len(name) > maxLength {
		return fmt.Errorf("Instance name must be between %d and %d characters", minLength, maxLength)
	}

	validNameRegexp := regexp.MustCompile(validRegexpString)

	matched := validNameRegexp.MatchString(name)
	if !matched {
		return errors.New("Invalid name: " + name)
	}

	return nil
}

// IsValidInput checks if the input value is in the range of valid values
func IsValidInput(input string, validValues ...string) bool {
	for _, b := range validValues {
		if input == b {
			return true
		}
	}

	return false
}
