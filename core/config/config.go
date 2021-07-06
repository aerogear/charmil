package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/pelletier/go-toml"
	"github.com/tidwall/sjson"
	"gopkg.in/yaml.v2"
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
		fileExt:  filepath.Ext(path),
	}

	return h
}

// Load reads config values from the local config file
// (using the file path linked to the handler) and stores
// them into the linked instance of host CLI config struct.
func (h *CfgHandler) Load() error {

	// Reads the local config file
	buf, err := readFile(h.filePath)
	if err != nil {
		return err
	}

	// Stores values (read from file) to the host config struct instance
	err = h.unmarshal(buf)
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
	buf, err := h.marshal()
	if err != nil {
		return err
	}

	// Writes the current host CLI config to the local config file
	err = writeFile(h.filePath, buf)
	if err != nil {
		return err
	}

	return nil
}

// MergePluginCfg adds config of specified plugin into the host CLI config struct.
func MergePluginCfg(pluginName string, h *CfgHandler, pluginCfg interface{}) error {

	// Stores the host CLI config struct as a byte array
	buf, err := h.marshal()
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
	err = h.unmarshal([]byte(mergedBuf))
	if err != nil {
		return err
	}

	return nil
}

// marshal identifies extension of the local config file and performs
// a marshaling operation on the host CLI config, based on it.
func (h *CfgHandler) marshal() ([]byte, error) {
	var marshalFunc func(in interface{}) ([]byte, error)

	switch h.fileExt {
	case ".yaml", ".yml":
		marshalFunc = yaml.Marshal

	case ".toml":
		marshalFunc = toml.Marshal

	case ".json":
		buf, err := json.MarshalIndent(h.cfg, "", "  ")
		if err != nil {
			return nil, err
		}
		return buf, nil

	default:
		return nil, fmt.Errorf("Unsupported file extension \"%v\"", h.fileExt)
	}

	buf, err := marshalFunc(h.cfg)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// unmarshal identifies extension of the local config file and performs
// an unmarshalling operation on the passed argument and the host CLI config
func (h *CfgHandler) unmarshal(in []byte) error {
	var unmarshalFunc func(in []byte, out interface{}) (err error)

	switch h.fileExt {
	case ".yaml", ".yml":
		unmarshalFunc = yaml.Unmarshal
	case ".json":
		unmarshalFunc = json.Unmarshal
	case ".toml":
		unmarshalFunc = toml.Unmarshal
	default:
		return fmt.Errorf("Unsupported file extension \"%v\"", h.fileExt)
	}

	err := unmarshalFunc(in, h.cfg)
	if err != nil {
		return err
	}

	return nil
}

// readFile reads the file specified by filePath and returns its contents.
func readFile(filePath string) ([]byte, error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// writeFile writes data to the file specified by filePath.
func writeFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}
