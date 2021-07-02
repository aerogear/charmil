package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// CfgFile defines the fields required to point to a local config file
type CfgFile struct {
	// Name of the config file (without extension)
	Name string

	// Extension of the config file.
	//
	// REQUIRED if the config file does not have the extension in the name
	Type string

	// Path to look for the config file in
	Path string
}

// Handler represents a wrapper around Viper
type Handler struct {
	vp *viper.Viper
}

// Plugin represents the config map imported from a plugin
type Plugin map[string]interface{}

// pluginCfg maps the name of plugins to their imported config maps
var pluginCfg = make(map[string]Plugin)

// New returns a new instance of the handler
func New() *Handler {
	h := &Handler{vp: viper.New()}
	return h
}

// InitFile links the handler instance to a local config file
// based on the specified file configuration
func (h *Handler) InitFile(f CfgFile) {
	h.vp.SetConfigName(f.Name)
	h.vp.SetConfigType(f.Type)
	h.vp.AddConfigPath(f.Path)
}

// Load imports config values from the local config file
// which was initialized while calling the InitFile method
func (h *Handler) Load() error {
	err := h.vp.ReadInConfig()
	return err
}

// Save writes the current config into the local config file
// which was initialized while calling the InitFile method
func (h *Handler) Save() error {
	err := h.vp.WriteConfig()
	return err
}

// SetValue stores the specified key, value pair in the current config
func (h *Handler) SetValue(key string, value interface{}) {
	h.vp.Set(key, value)
}

// GetValue returns value of the specified key from the current config
func (h *Handler) GetValue(key string) (interface{}, error) {
	if h.vp.IsSet(key) {
		return h.vp.Get(key), nil
	}
	return nil, fmt.Errorf("Key doesn't exist")
}

// GetAllSettings returns the current config in the form of a map
func (h *Handler) GetAllSettings() map[string]interface{} {
	return h.vp.AllSettings()
}

// SetPluginCfg stores the imported plugin config as a key, value pair in a map,
// where the key represents name of the plugin and the value being its config map.
//
// For eg. Key: "pluginA", Value: map[key5:value5 key6:value6 key7:value7 key8:value8]
func (h *Handler) SetPluginCfg(pluginName string, p Plugin) {
	pluginCfg[pluginName] = p
}

// MergePluginCfg stores the pluginCfg map into the current config
func (h *Handler) MergePluginCfg() {
	h.SetValue("plugins", pluginCfg)
}
