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
	language  *language.Tag
	bundle    *i18n.Bundle
	Localizer *i18n.Localizer
	format    string
	path      string
}

type Config struct {
	language language.Tag
	path     string
	format   string
}

func (I *GoI18n) LocalizeByID(messageId string, templateData map[string]interface{}) (string, error) {

	localizeConfig := &i18n.LocalizeConfig{MessageID: messageId, TemplateData: templateData}
	res := I.Localizer.MustLocalize(localizeConfig)

	return res, nil
}

func (I *GoI18n) InitLocalizer(cfg Config) (*GoI18n, error) {
	bundle := i18n.NewBundle(cfg.language)

	var unmarshalFunc i18n.UnmarshalFunc

	switch cfg.format {
	case "toml":
		unmarshalFunc = toml.Unmarshal
	case "json":
		unmarshalFunc = json.Unmarshal
	case "yaml":
		unmarshalFunc = yaml.Unmarshal
	default:
		return nil, errors.New("unsupported format " + cfg.format)
	}

	bundle.RegisterUnmarshalFunc(cfg.format, unmarshalFunc)
	bundle.LoadMessageFile(cfg.path)

	localizer := i18n.NewLocalizer(bundle)

	res := &GoI18n{
		language:  &cfg.language,
		bundle:    bundle,
		format:    cfg.format,
		path:      cfg.path,
		Localizer: localizer,
	}

	return res, nil
}
