package factory

import (
	"fmt"
	"os"

	"github.com/aerogear/charmil/core/localize"
	"github.com/aerogear/charmil/core/logging"
)

// Default creates new instance of factory for plugins
func Default(localizer localize.Localizer) *Factory {

	// initializing logger
	loggerFunc := func() (logging.Logger, error) {
		loggerBuilder := logging.NewStdLoggerBuilder()
		loggerBuilder = loggerBuilder.Streams(os.Stdout, os.Stdin)
		logger, err := loggerBuilder.Build()
		if err != nil {
			return nil, err
		}

		return logger, nil
	}

	logger, err := loggerFunc()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return &Factory{
		Logger:    logger,
		Localizer: localizer,
	}
}
