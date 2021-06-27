package rules

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	ErrMin           = errors.New("less than min")
	ErrMax           = errors.New("less than max")
	ErrLen           = errors.New("invalid length")
	ErrNeg           = errors.New("negative value")
	ErrZeroValue     = errors.New("zero value")
	ErrFieldNotExist = errors.New("field doesn't exists")
	ErrRegexp        = errors.New("regular expression mismatch")
	ErrUnsupported   = errors.New("unsupported type")
	ErrAbsent        = errors.New("field must be present")
)

// ValidationError is a default validation error
type ValidationError struct {
	Name string
	Err  error
	Rule string
	cmd  *cobra.Command
}

// StatusLog is used for providing info
// about validation of commands
type StatusLog struct {
	// totalTested represents number of commands checked
	totalTested int
	// totalErrors reperesents number of errors
	totalErrors int
	// errors in command
	errors []ValidationError
}
