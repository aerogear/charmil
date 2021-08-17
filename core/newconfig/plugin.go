package newconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// TODO: Remove the hardcoded name and get this value dynamically
const pluginName = "plugin1"

type pluginCfgHandler struct{}

func (p *pluginCfgHandler) Load() (*Config, error) {
	filePath, err := p.Location()
	if err != nil {
		return nil, err
	}

	data, err := readFile(filePath)
	if err != nil {
		return nil, err
	}

	subVal := getJsonSubkey(data, "Plugins."+pluginName)

	var cfg Config

	err = json.Unmarshal(subVal, &cfg)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", "unable to parse config", err)
	}

	return &cfg, nil
}

func (p *pluginCfgHandler) Save(cfg *Config) error {
	f, err := p.Location()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("%v: %w", "unable to marshal config", err)
	}

	finalData, err := setJsonSubkey(data, "Plugins."+pluginName, cfg)
	if err != nil {
		return err
	}

	err = writeFile(f, finalData)
	if err != nil {
		return err
	}

	return nil
}

func (p *pluginCfgHandler) Remove() error {
	filePath, err := p.Location()
	if err != nil {
		return err
	}

	data, err := readFile(filePath)
	if err != nil {
		return err
	}

	finalData, err := deleteJsonSubkey(data, "Plugins."+pluginName)
	if err != nil {
		return err
	}

	err = writeFile(filePath, finalData)
	if err != nil {
		return err
	}

	return nil
}

func (p *pluginCfgHandler) Location() (path string, err error) {
	if hasCustomLocation() {
		path = os.Getenv(envName)
	} else {
		cfgDir, err := defaultDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(cfgDir, "config.json")
		if err != nil {
			return "", err
		}
	}

	return path, nil
}
