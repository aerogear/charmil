package factory

import (
	"os"

	"github.com/aerogear/charmil/pkg/localize"
	"github.com/aerogear/charmil/pkg/logging"
)

// create new instance of factory for plugins
func Default(localizer localize.Localizer) *Factory {
	var logger logging.Logger

	loggerFunc := func() (logging.Logger, error) {
		if logger != nil {
			return logger, nil
		}

		loggerBuilder := logging.NewStdLoggerBuilder()
		loggerBuilder = loggerBuilder.Streams(os.Stdout, os.Stdin)
		logger, err := loggerBuilder.Build()

		if err != nil {
			return nil, err
		}

		return logger, nil
	}

	return &Factory{
		Logger:    loggerFunc,
		Localizer: localizer,
	}
}
