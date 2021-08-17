package newconfig

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type hostCfgHandler struct{}

func (h *hostCfgHandler) Load() (*Config, error) {
	filePath, err := h.Location()
	if err != nil {
		return nil, err
	}

	data, err := readFile(filePath)
	if err != nil {
		return nil, err
	}

	var cfg Config

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", "unable to parse config", err)
	}

	return &cfg, nil
}

func (h *hostCfgHandler) Save(cfg *Config) error {
	f, err := h.Location()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("%v: %w", "unable to marshal config", err)
	}

	err = writeFile(f, data)
	if err != nil {
		return err
	}

	return nil
}

func (h *hostCfgHandler) Remove() error {
	f, err := h.Location()
	if err != nil {
		return err
	}

	_, err = os.Stat(f)
	if os.IsNotExist(err) {
		return nil
	}

	err = os.Remove(f)
	if err != nil {
		return err
	}

	return nil
}

func (h *hostCfgHandler) Location() (path string, err error) {
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
