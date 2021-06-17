package localize

type Localizer interface {
	LocalizeByID(messageId string) string
}
