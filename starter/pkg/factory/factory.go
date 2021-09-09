package factory

import (
	"github.com/aerogear/charmil/core/config"
	"github.com/aerogear/charmil/core/utils/iostreams"
	"github.com/aerogear/charmil/core/utils/localize"
	"github.com/aerogear/charmil/core/utils/logging"
)

// Factory is an abstract type which provides
// the access of charmil packages to the commands
type Factory struct {
	// Type which defines the streams for the CLI
	IOStreams *iostreams.IOStreams
	// Logger provides functions for unified logging
	Logger logging.Logger
	// Localizer localizes the text in commands
	Localizer localize.Localizer
	// CfgHandler provides the fields required for managing config
	CfgHandler *config.CfgHandler
}
