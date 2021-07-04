package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"

	toml "github.com/pelletier/go-toml"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
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
	vp         *viper.Viper
	cfg        interface{}
	fileFormat string
}

// New returns a new instance of the handler
func New(f CfgFile, cfg interface{}) *Handler {
	// TODO: Add code to check if cfg is of type: struct

	// fmt.Println(reflect.TypeOf(cfg))
	// t := reflect.ValueOf(cfg).Kind()
	// if t != reflect.Struct {
	// 	return fmt.Errorf("The object passed in as argument is not a struct.", t)
	// }

	h := &Handler{vp: viper.New(), cfg: cfg, fileFormat: f.Type}

	h.vp.SetConfigName(f.Name)
	h.vp.SetConfigType(f.Type)
	h.vp.AddConfigPath(f.Path)

	return h
}

func (h *Handler) unmarshal() error {
	err := h.vp.Unmarshal(h.cfg)
	if err != nil {
		return err
	}

	return nil
}

func (h *Handler) Save() error {
	bs, err := h.marshal()
	if err != nil {
		log.Fatal(err)
	}

	h.readConfig(bs)

	err = h.vp.WriteConfig()

	return err
}

// Load imports config values from the local config file
// which was initialized while calling the InitFile method
func (h *Handler) Load() error {
	err := h.vp.ReadInConfig()

	err = h.unmarshal()
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (h *Handler) readConfig(buf []byte) {
	h.vp.ReadConfig(bytes.NewBuffer(buf))
}

func (h *Handler) marshal() ([]byte, error) {
	var marshalFunc func(v interface{}) ([]byte, error)

	switch h.fileFormat {
	case "yaml", "yml":
		marshalFunc = yaml.Marshal
	case "json":
		marshalFunc = json.Marshal
	case "toml":
		marshalFunc = toml.Marshal
	default:
		return nil, fmt.Errorf("Unsupported format \"%v\"", h.fileFormat)
	}

	bs, err := marshalFunc(h.cfg)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (h *Handler) SavePluginCfg() error {
	bs, err := h.marshal()
	if err != nil {
		log.Fatal(err)
	}

	h.readConfig(bs)

	err = h.vp.WriteConfig()

	return err
}

// func (h *Handler) LoadPluginCfg() error {
// 	err := h.vp.ReadInConfig()

// 	err = h.unmarshal()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return err
// }
