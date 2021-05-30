package pluginloader

type PluginConfig struct {
	Commands []CommandConfig `yaml:"commands"`
}

type CommandConfig struct {
	Name             string       `yaml:"name"`
	MapsTo           ArgsConfig   `yaml:"mapsTo"`
	Flags            []FlagConfig `yaml:"flags"`
	ShortDescription string       `yaml:"shortDescription"`
	Examples         string       `yaml:"usage"`
}

type ArgsConfig struct {
	Name       string
	Subcommand string
	Args       []string
}

type FlagConfig struct {
	Type         string `yaml:"type"`
	DefaultValue string `yaml:"defaultValue"`
	Name         string `yaml:"name"`
	MapsTo       string `yaml:"mapsTo"`
	Description  string `yaml:"description"`
	Alias        string `yaml:"alias"`
}
