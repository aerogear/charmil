package localize

// Interface for localizing messages from locals
type Localizer interface {
	// localize by providing id of text in locals
	LocalizeByID(messageId string) string
}
