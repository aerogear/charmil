package factory

import (
	"github.com/aerogear/charmil/pkg/logging"
)

// Access to all the commands
type Factory struct {
	Logger func() (logging.Logger, error)
}
