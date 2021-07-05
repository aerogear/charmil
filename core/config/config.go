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

type Handler struct {
	cfg      interface{}
	filePath string
	fileExt  string
}

func New(path string, cfg interface{}) *Handler {
	// TODO: Add code to check if cfg is of type: struct

	// fmt.Println(reflect.TypeOf(cfg))
	// t := reflect.ValueOf(cfg).Kind()
	// if t != reflect.Struct {
	// 	return fmt.Errorf("The object passed in as argument is not a struct.", t)
	// }

	h := &Handler{
		filePath: path,
		cfg:      cfg,
		fileExt:  filepath.Ext(path),
	}

	return h
}

func (h *Handler) Load() error {
	// Add code to read local config file

	buf, err := readFile(h.filePath)
	if err != nil {
		return err
	}

	// Add code to store values (read from file) to the host config struct [ie. Unmarshal]

	err = h.unmarshal(buf, h.cfg)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) Save() error {
	// Add code to store values (read from host config struct) to the local config file [ie. Marshal]

	bs, err := h.marshal()
	if err != nil {
		return err
	}

	err = writeFile(h.filePath, bs)
	if err != nil {
		return err
	}

	return nil
}

func MergePluginCfg(pluginName string, cfgFilePath string, cfg interface{}) error {
	// TODO: Add code to check if cfg is of type: struct

	// Load local config file content into a byte-array/string [Marshal]

	bs, err := readFile(cfgFilePath)
	if err != nil {
		return err
	}

	// TODO: Specific to JSON files currently. Extend to other formats too
	updatedBs, err := sjson.Set(string(bs), "plugins."+pluginName, cfg)
	if err != nil {
		return err
	}

	// Write final string to the local config file
	err = writeFile(cfgFilePath, []byte(updatedBs))
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) marshal() ([]byte, error) {
	var marshalFunc func(in interface{}) ([]byte, error)

	switch h.fileExt {
	case ".yaml", ".yml":
		marshalFunc = yaml.Marshal

	case ".toml":
		marshalFunc = toml.Marshal

	case ".json":
		bs, err := json.MarshalIndent(h.cfg, "", "  ")
		if err != nil {
			return nil, err
		}
		return bs, nil

	default:
		return nil, fmt.Errorf("Unsupported file extension \"%v\"", h.fileExt)
	}

	bs, err := marshalFunc(h.cfg)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (h *Handler) unmarshal(in []byte, out interface{}) error {
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

func readFile(filePath string) ([]byte, error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func writeFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
