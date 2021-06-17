package factory

import (
	"fmt"
	"os"

	"github.com/aerogear/charmil/pkg/localize"
	"github.com/aerogear/charmil/pkg/logging"
	"golang.org/x/text/language"
)

// create new command factory
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

	loc, err := localize.InitLocalizer(localize.Config{Language: language.English, Path: "examples/plugins/date/en.yaml", Format: "yaml"})

	if err != nil {
		fmt.Println("Error", err)
	}

	return &Factory{
		Logger:    loggerFunc,
		Localizer: loc,
	}
}
