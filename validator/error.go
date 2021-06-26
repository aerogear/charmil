package validator

import "errors"

var (
	ErrMin           = errors.New("less than min")
	ErrMax           = errors.New("less than max")
	ErrLen           = errors.New("invalid length")
	ErrNeg           = errors.New("negative value")
	ErrRegexp        = errors.New("regular expression mismatch")
	ErrUnsupported   = errors.New("unsupported type")
	ErrZeroValue     = errors.New("zero value")
	ErrAbsent        = errors.New("field must be present")
	ErrFieldNotExist = errors.New("field doesn't exists")
)

type Error struct {
	Name string
	Err  error
	Rule string
}
