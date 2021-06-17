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
	Language language.Tag
	Path     string
	Format   string
}

func (I *GoI18n) LocalizeByID(messageId string) string {

	// localizeConfig := &i18n.LocalizeConfig{MessageID: messageId, TemplateData: templateData}

	localizeConfig := &i18n.LocalizeConfig{MessageID: messageId, PluralCount: 1}
	res := I.Localizer.MustLocalize(localizeConfig)

	return res
}

func InitLocalizer(cfg Config) (*GoI18n, error) {
	bundle := i18n.NewBundle(cfg.Language)

	var unmarshalFunc i18n.UnmarshalFunc

	switch cfg.Format {
	case "toml":
		unmarshalFunc = toml.Unmarshal
	case "json":
		unmarshalFunc = json.Unmarshal
	case "yaml":
		unmarshalFunc = yaml.Unmarshal
	default:
		return nil, errors.New("unsupported format " + cfg.Format)
	}

	bundle.RegisterUnmarshalFunc(cfg.Format, unmarshalFunc)
	bundle.LoadMessageFile(cfg.Path)

	loc := &GoI18n{
		language:  &cfg.Language,
		bundle:    bundle,
		format:    cfg.Format,
		path:      cfg.Path,
		Localizer: i18n.NewLocalizer(bundle),
	}

	return loc, nil
}
