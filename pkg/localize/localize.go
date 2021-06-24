package localize

// Localizer is an Interface for localizing messages from locals
type Localizer interface {
	// LocalizeById takes ID
	LocalizeByID(messageId string, template ...*TemplateEntry) string
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
