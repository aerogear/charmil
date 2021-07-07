package localize

import (
	"embed"
	"io/fs"

	"golang.org/x/text/language"
)

var (
	//go:embed locales
	defaultLocales  embed.FS
	defaultLanguage = &language.English
)

// Localizer is an abstract interface
// which defines methods to load i18n messages
type Localizer interface {
	MustLocalize(id string, templateEntries ...*TemplateEntry) string
}

// TemplateEntry is a type which defines
// variable interpolation key:value pairs
// which are used to pass dynamic values to the template
type TemplateEntry struct {
	Key   string
	Value interface{}
}

// NewEntry returns a new template entry
// which is a type for interpolating a string
func NewEntry(key string, val interface{}) *TemplateEntry {
	return &TemplateEntry{key, val}
}

// GetDefaultLocales returns the default locale files
func GetDefaultLocales() fs.FS {
	return defaultLocales
}

// GetDefaultLanguage returns the default i18n language used
func GetDefaultLanguage() *language.Tag {
	return defaultLanguage
}
