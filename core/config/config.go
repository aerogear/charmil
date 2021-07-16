package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		return err
	}

	// Stores values (read from file) to the host config struct instance
	err = Unmarshal(buf, h.cfg, h.fileExt)
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
	buf, err := Marshal(h.cfg, h.fileExt)
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
	buf, err := Marshal(h.cfg, h.fileExt)
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
	err = Unmarshal([]byte(mergedBuf), h.cfg, h.fileExt)
	if err != nil {
		return err
	}

	return nil
}

// Marshal converts the passed object into byte data, based on the specified file format
func Marshal(in interface{}, fileExt string) ([]byte, error) {
	var marshalFunc func(in interface{}) ([]byte, error)

	switch fileExt {
	case ".yaml", ".yml":
		marshalFunc = yaml.Marshal

	case ".toml":
		marshalFunc = toml.Marshal

	case ".json":
		buf, err := json.MarshalIndent(in, "", "  ")
		if err != nil {
			return nil, err
		}
		return buf, nil

	default:
		return nil, fmt.Errorf("Unsupported file extension \"%v\"", fileExt)
	}

	buf, err := marshalFunc(in)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// Unmarshal converts the passed byte data into a struct
func Unmarshal(in []byte, out interface{}, fileExt string) error {
	var unmarshalFunc func(in []byte, out interface{}) (err error)

	switch fileExt {
	case ".yaml", ".yml":
		unmarshalFunc = yaml.Unmarshal
	case ".json":
		unmarshalFunc = json.Unmarshal
	case ".toml":
		unmarshalFunc = toml.Unmarshal
	default:
		return fmt.Errorf("Unsupported file extension \"%v\"", fileExt)
	}

	err := unmarshalFunc(in, out)
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
