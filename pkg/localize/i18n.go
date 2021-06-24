package localize

import (
	"encoding/json"
	"errors"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

// go-i18n implementation

// GoI18n is a type which
// stores the details to create bundle
type GoI18n struct {
	language  *language.Tag
	bundle    *i18n.Bundle
	Localizer *i18n.Localizer
	format    string
	path      string
}

// Config is a type which helps to get the
// information to initialize the localizer
type Config struct {
	Language language.Tag
	Path     string
	Format   string
}

// LocalizeByID helps in localizing by passing id present in local file
// pass dynamic value using template entries
func (I *GoI18n) LocalizeByID(messageId string, template ...*TemplateEntry) string {

	// Putting back templateEntry into desired format
	// required by go-i18n
	templateData := map[string]interface{}{}
	for _, t := range template {
		templateData[t.Key] = t.Value
	}

	localizeConfig := &i18n.LocalizeConfig{MessageID: messageId, PluralCount: 1, TemplateData: templateData}
	res := I.Localizer.MustLocalize(localizeConfig)
	return res
}

// InitLocalizer initialize the localizer
// and create new instance for plugin
// Pass Config including lang, path & format of locals
func InitLocalizer(cfg Config) (*GoI18n, error) {

	// create bundle of choose language
	bundle := i18n.NewBundle(cfg.Language)
	var unmarshalFunc i18n.UnmarshalFunc

	// choose unmarshal func according to format of local
	switch cfg.Format {
	case "toml":
		unmarshalFunc = toml.Unmarshal
	case "json":
		unmarshalFunc = json.Unmarshal
	case "yaml":
		unmarshalFunc = yaml.Unmarshal
	default:
		return nil, errors.New("unsupported format of local file " + cfg.Format)
	}

	// load translations during initialization
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
