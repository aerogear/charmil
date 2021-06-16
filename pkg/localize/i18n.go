package localize

import (
	"encoding/json"
	"errors"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

type GoI18n struct {
}

type Config struct {
	language language.Tag
	path     string
	format   string
}

func (I *GoI18n) Localize(confg Config, messageId string, templateData map[string]interface{}) (string, error) {
	bundle := i18n.NewBundle(confg.language)

	var unmarshalFunc i18n.UnmarshalFunc

	switch confg.format {
	case "toml":
		unmarshalFunc = toml.Unmarshal
	case "json":
		unmarshalFunc = json.Unmarshal
	case "yaml":
		unmarshalFunc = yaml.Unmarshal
	default:
		return "", errors.New("unsupported format " + confg.format)
	}

	bundle.RegisterUnmarshalFunc(confg.format, unmarshalFunc)
	bundle.LoadMessageFile(confg.path)

	localizer := i18n.NewLocalizer(bundle)

	localizeConfig := &i18n.LocalizeConfig{MessageID: messageId, TemplateData: templateData}
	res := localizer.MustLocalize(localizeConfig)

	return res, nil
}
