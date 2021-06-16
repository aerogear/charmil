package localize

type Localizer interface {
	LocalizeByID(confg Config, messageId string, templateData map[string]interface{})
}
