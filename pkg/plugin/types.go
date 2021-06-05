package core

type PluginConfig struct {
	Commands []CommandConfig
}

// CommandConfig
type CommandConfig struct {
	Name             string
	MapsTo           ArgsMapConfig
	Flags            []FlagConfig
	ShortDescription string
	Examples         string
}

type ArgsMapConfig struct {
	Name       string
	Subcommand string
	Args       []string
}

// FlagConfig
type FlagConfig struct {
	Type         string
	DefaultValue string
	Name         string
	MapsTo       string
	Description  string
	Alias        string
}
