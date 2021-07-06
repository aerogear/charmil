package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// NewFile creates a new config type
func NewFile() IConfig {
	cfg := &File{}

	return cfg
}

// File is a type which describes a config file
type File struct{}

const errorFormat = "%v: %w"

// Load loads the configuration from the configuration file. If the configuration file doesn't exist
// it will return an empty configuration object.
func (c *File) Load() (*Config, error) {
	file, err := c.Location()
	if err != nil {
		return nil, err
	}
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf(errorFormat, "unable to check if config file exists", err)
	}
	// #nosec G304
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf(errorFormat, "unable to read config file", err)
	}
	var cfg Config
	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return nil, fmt.Errorf(errorFormat, "unable to parse config", err)
	}
	return &cfg, nil
}

// Save saves the given configuration to the configuration file.
func (c *File) Save(cfg *Config) error {
	file, err := c.Location()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return fmt.Errorf("%v: %w", "unable to marshal config", err)
	}
	cfgDir, err := DefaultDir()
	if err != nil {
		return err
	}
	if _, err = os.Stat(cfgDir); os.IsNotExist(err) {
		err = os.Mkdir(cfgDir, 0o700)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(file, data, 0o600)
	if err != nil {
		return fmt.Errorf(errorFormat, "unable to save config", err)
	}
	return nil
}

// Remove removes the configuration file.
func (c *File) Remove() error {
	file, err := c.Location()
	if err != nil {
		return err
	}
	_, err = os.Stat(file)
	if os.IsNotExist(err) {
		return nil
	}
	err = os.Remove(file)
	if err != nil {
		return err
	}
	return nil
}

// Location gets the path to the config file
func (c *File) Location() (path string, err error) {
	if cfgDir := os.Getenv("CHARMIL_CLI_CONFIG"); cfgDir != "" {
		path = cfgDir
	} else {
		cfgDir, err := DefaultDir()
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

// DefaultDir returns the default parent directory of the config file
func DefaultDir() (string, error) {
	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(userCfgDir, "charmil-starter"), nil
}
