package io

import (
	"encoding/json"
	"fmt"
	"github.com/pelletier/go-toml"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

// Marshal reads bytes from the given file path
// in the form of the given template
func Marshal(template *interface{}, path string) ([]byte, error) {
	var marshalFunc func(in interface{}) ([]byte, error)

	extension := filepath.Ext(path)
	switch extension {
	case ".yaml", ".yml":
		marshalFunc = yaml.Marshal
	case ".toml":
		marshalFunc = toml.Marshal
	case ".json":
		buf, err := json.MarshalIndent(template, "", "  ")
		if err != nil {
			return nil, err
		}
		return buf, nil
	default:
		return nil, fmt.Errorf("Unsupported file extension \"%v\" ", extension)
	}

	buf, err := marshalFunc(template)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// Unmarshal populates the receiver with the given bytes
// based on the given file those bytes are from
func Unmarshal(in []byte, receiver *interface{}, path string) error {
	var unmarshalFunc func(in []byte, out interface{}) (err error)

	extension := filepath.Ext(path)
	switch extension {
	case ".yaml", ".yml":
		unmarshalFunc = yaml.Unmarshal
	case ".json":
		unmarshalFunc = json.Unmarshal
	case ".toml":
		unmarshalFunc = toml.Unmarshal
	default:
		return fmt.Errorf("Unsupported file extension \"%v\" ", extension)
	}

	err := unmarshalFunc(in, *receiver)
	if err != nil {
		return err
	}

	return nil
}

// ReadFile reads the file specified by filePath and returns its contents.
func ReadFile(filePath string) ([]byte, error) {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

// WriteFile writes data to the file specified by filePath.
func WriteFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}