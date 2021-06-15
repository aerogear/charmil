package factory

import (
	"os"

	"github.com/redhat-developer/app-services-cli/pkg/logging"
)

type Factory struct {
	// Returns a logger to create leveled logs in the application
	Logger func() (logging.Logger, error)
}

func Default() *Factory {
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
		Logger: loggerFunc,
	}
}
