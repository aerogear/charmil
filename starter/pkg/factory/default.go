package factory

import (
	"fmt"

	"github.com/aerogear/charmil/core/config"
	"github.com/aerogear/charmil/core/utils/iostreams"
	"github.com/aerogear/charmil/core/utils/localize"
	"github.com/aerogear/charmil/core/utils/logging"
)

// Default creates new instance of factory for plugins
func Default(localizer localize.Localizer, cfgHandler *config.CfgHandler) *Factory {
	io := iostreams.System()

	var logger logging.Logger
	// initializing logger
	loggerFunc := func() (logging.Logger, error) {
		if logger != nil {
			return logger, nil
		}

		loggerBuilder := logging.NewStdLoggerBuilder()
		loggerBuilder = loggerBuilder.Streams(io.Out, io.ErrOut)

		debugEnabled := true
		loggerBuilder = loggerBuilder.Debug(debugEnabled)

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
		IOStreams:  io,
		Logger:     logger,
		Localizer:  localizer,
		CfgHandler: cfgHandler,
	}
}
