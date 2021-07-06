package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/imdario/mergo"
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
	// TODO: Add code to verify if cfg is a pointer to struct

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
	err = h.unmarshal(buf, h.cfg)
	if err != nil {
		return err
	}

	return nil
}

// Save writes config values from the linked instance
// of host CLI config struct to the local config file
// (using the file path linked to the handler).
func (h *CfgHandler) Save() error {
	// To store the current contents of the local config file
	dst := &map[string]interface{}{}

	// To store contents of the linked instance of host CLI config struct
	src := &map[string]interface{}{}

	// Reads the local config file
	buf, err := readFile(h.filePath)
	if err != nil {
		return err
	}

	// Stores current contents of the local config file to `dst`
	err = h.unmarshal(buf, &dst)
	if err != nil {
		return err
	}

	// Initializes a `map[string]interface{}` holding the values of `h.cfg`.
	//
	// Done with the intention to maintain similar types as `mergo.Merge`
	// doesn't accept 2 differently typed objects as arguments
	err = mergo.Map(src, h.cfg)
	if err != nil {
		return err
	}

	// Merges current host CLI config with the existing config in the local file
	if err = mergo.Merge(dst, src, mergo.WithSliceDeepCopy); err != nil {
		return err
	}

	// Converts the merged config into a byte array
	bs, err := h.marshal(*dst)
	if err != nil {
		return err
	}

	// Writes the merged config to the local config file
	err = writeFile(h.filePath, bs)
	if err != nil {
		return err
	}

	return nil
}

// MergePluginCfg adds plugin config into the local config file.
func MergePluginCfg(pluginName string, h *CfgHandler, pluginCfg interface{}) error {
	// TODO: Add code to verify if `pluginCfg` is a pointer to struct

	// Stores the host CLI config struct as a byte array
	buf, err := h.marshal(h.cfg)
	if err != nil {
		return err
	}

	// TODO: Specific only to JSON files currently. Extend to other formats too

	// Adds a field (specified by `pluginName`) under the `Plugins` field of
	// host CLI config struct and stores the specified plugin config under that sub-field
	mergedBuf, err := sjson.Set(string(buf), "Plugins."+pluginName, pluginCfg)
	if err != nil {
		return err
	}

	// Updates the host CLI config struct with merged plugin config
	err = h.unmarshal([]byte(mergedBuf), h.cfg)
	if err != nil {
		return err
	}

	return nil
}

// marshal identifies extension of the local config file and performs
// a marshaling operation on the passed object, based on it.
func (h *CfgHandler) marshal(in interface{}) ([]byte, error) {
	var marshalFunc func(in interface{}) ([]byte, error)

	switch h.fileExt {
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
		return nil, fmt.Errorf("Unsupported file extension \"%v\"", h.fileExt)
	}

	buf, err := marshalFunc(in)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// unmarshal identifies extension of the local config file and performs
// an unmarshalling operation on the passed arguments, based on it.
func (h *CfgHandler) unmarshal(in []byte, out interface{}) error {
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
