package pluginloader

type CommandConfig struct {
	Name             string
	Args             []string
	Flags            []FlagConfig
	ShortDescription string
	Examples         string
}

type FlagConfig struct {
	Type         string
	DefaultValue string
	Name         string
	Description  string
	Alias        string
}
