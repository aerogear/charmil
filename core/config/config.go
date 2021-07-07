package config

import (
	"github.com/aerogear/charmil/core/io"
	"github.com/tidwall/sjson"
)

// CfgHandler defines the fields required to manage config.
type CfgHandler struct {
	// Pointer to an instance of the host CLI config struct
	cfg interface{}

	// Path of the local config file
	filePath string

	// Extension of the local config file
	fileExt string
}

// NewHandler links the specified arguments to a
// new instance of config handler and returns a pointer to it.
func NewHandler(path string, cfg interface{}) *CfgHandler {
	h := &CfgHandler{
		filePath: path,
		cfg:      cfg,
	}

	return h
}

// Load reads config values from the local config file
// (using the file path linked to the handler) and stores
// them into the linked instance of host CLI config struct.
func (h *CfgHandler) Load() error {

	// Reads the local config file
	buf, err := io.ReadFile(h.filePath)
	if err != nil {
		return err
	}

	// Stores values (read from file) to the host config struct instance
	err = io.Unmarshal(buf, &h.cfg, h.filePath)
	if err != nil {
		return err
	}

	return nil
}

// Save writes config values from the linked instance
// of host CLI config struct to the local config file
// (using the file path linked to the handler).
func (h *CfgHandler) Save() error {
	// Stores the host CLI config as a byte array
	buf, err := io.Marshal(&h.cfg, h.filePath)
	if err != nil {
		return err
	}

	// Writes the current host CLI config to the local config file
	err = io.WriteFile(h.filePath, buf)
	if err != nil {
		return err
	}

	return nil
}

// MergePluginCfg adds config of specified plugin into the host CLI config struct.
func MergePluginCfg(pluginName string, h *CfgHandler, pluginCfg interface{}) error {

	// Stores the host CLI config struct as a byte array
	buf, err := io.Marshal(&h.cfg, h.filePath)
	if err != nil {
		return err
	}

	// Adds a field (specified by `pluginName`) under the `Plugins` field of
	// host CLI config struct and stores the specified plugin config under that sub-field
	mergedBuf, err := sjson.Set(string(buf), "Plugins."+pluginName, pluginCfg)
	if err != nil {
		return err
	}

	// Updates the host CLI config struct with merged plugin config
	err = io.Unmarshal([]byte(mergedBuf), &h.cfg, h.filePath)
	if err != nil {
		return err
	}

	return nil
}
