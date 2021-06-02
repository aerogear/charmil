package pluginloader

type PluginConfig struct {
	Commands []CommandConfig
}

type CommandConfig struct {
	Name             string
	MapsTo           ArgsConfig
	Flags            []FlagConfig
	ShortDescription string
	Examples         string
}

type ArgsConfig struct {
	Name       string
	Subcommand string
	Args       []string
}

type FlagConfig struct {
	Type         string
	DefaultValue string
	Name         string
	MapsTo       string
	Description  string
	Alias        string
}
