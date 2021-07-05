package factory

import (
	"github.com/aerogear/charmil/core/localize"
	"github.com/aerogear/charmil/core/logging"
)

// Factory is an abstract type which provides
// the access of charmil packages to the commands
type Factory struct {
	// Logger provides functions for unified logging
	Logger logging.Logger
	// Localizer localizes the text in commands
	Localizer localize.Localizer
}
