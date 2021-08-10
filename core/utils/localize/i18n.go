package localize

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"

	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

// go-i18n implementation

// GoI18n is a type which
// stores the details to create bundle
type GoI18n struct {
	files     fs.FS
	language  *language.Tag
	bundle    *i18n.Bundle
	Localizer *i18n.Localizer
	format    string
	path      string
}

// Config is a type which helps to get the
// information to initialize the localizer
type Config struct {
	Files    embed.FS
	Language *language.Tag
	Path     string
	Format   string
}

// LocalizeByID helps in localizing by passing id present in local file
// pass dynamic value using template entries
func (i *GoI18n) LocalizeByID(messageId string, template ...*TemplateEntry) string {

	// Putting back templateEntry into desired format
	// required by go-i18n
	templateData := map[string]interface{}{}
	for _, t := range template {
		templateData[t.Key] = t.Value
	}

	localizeConfig := &i18n.LocalizeConfig{MessageID: messageId, PluralCount: 1, TemplateData: templateData}
	res := i.Localizer.MustLocalize(localizeConfig)
	return res
}

func New(cfg *Config) (Localizer, error) {
	if cfg == nil {
		cfg = &Config{}
	}

	if cfg.Format == "" {
		cfg.Format = "toml"
	}

	bundle := i18n.NewBundle(*cfg.Language)
	loc := &GoI18n{
		files:     cfg.Files,
		language:  cfg.Language,
		bundle:    bundle,
		Localizer: i18n.NewLocalizer(bundle),
		format:    cfg.Format,
		path:      cfg.Path,
	}

	err := loc.load()
	return loc, err
}

// walk the file system and load each file into memory
func (i *GoI18n) load() error {
	// localesRoot := filepath.Join(i.path, i.language.String())
	return fs.WalkDir(i.files, ".", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		return i.MustLocalizeFile(i.files, path)
	})
}

// read the message file from the file system
func (i *GoI18n) MustLocalizeFile(files fs.FS, path string) (err error) {
	// open the static i18n file
	buf, err := fs.ReadFile(files, path)
	if err != nil {
		return err
	}
	fileext := fmt.Sprintf("%v.%v", i.language.String(), i.format)
	var unmarshalFunc i18n.UnmarshalFunc
	switch i.format {
	case "toml":
		unmarshalFunc = toml.Unmarshal
	case "yaml", "yml":
		unmarshalFunc = yaml.Unmarshal
	case "json":
		unmarshalFunc = json.Unmarshal
	default:
		return fmt.Errorf("unsupported format \"%v\"", i.format)
	}

	i.bundle.RegisterUnmarshalFunc(i.format, unmarshalFunc)
	_, err = i.bundle.ParseMessageFileBytes(buf, fileext)

	return
}
