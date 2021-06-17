package factory

import (
	"github.com/aerogear/charmil/pkg/localize"
	"github.com/aerogear/charmil/pkg/logging"
)

// Access to all the commands
type Factory struct {
	// Logger provides functions for unified logging
	Logger func() (logging.Logger, error)
	// Localizer localizes the text in commands
	Localizer localize.Localizer
}
