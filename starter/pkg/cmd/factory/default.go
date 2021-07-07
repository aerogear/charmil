package factory

import (
	"github.com/aerogear/charmil/internal/config"
	"github.com/aerogear/charmil/pkg/cmd/debug"
	"github.com/aerogear/charmil/pkg/iostreams"
	"github.com/aerogear/charmil/pkg/localize"
	"github.com/aerogear/charmil/pkg/logging"
)

// New creates a new command factory
// The command factory is available to all command packages
// giving centralized access to the config

// nolint:funlen
func New(cliVersion string, localizer localize.Localizer) *Factory {
	io := iostreams.System()

	var logger logging.Logger
	cfgFile := config.NewFile()

	loggerFunc := func() (logging.Logger, error) {
		if logger != nil {
			return logger, nil
		}

		loggerBuilder := logging.NewStdLoggerBuilder()
		loggerBuilder = loggerBuilder.Streams(io.Out, io.ErrOut)

		debugEnabled := debug.Enabled()
		loggerBuilder = loggerBuilder.Debug(debugEnabled)

		logger, err := loggerBuilder.Build()
		if err != nil {
			return nil, err
		}

		return logger, nil
	}

	return &Factory{
		IOStreams: io,
		Config:    cfgFile,
		Logger:    loggerFunc,
		Localizer: localizer,
	}
}
