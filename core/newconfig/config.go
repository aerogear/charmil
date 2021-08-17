package newconfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

const envName = "CHARMIL_CONFIG"

// TODO: Temporary struct. Find a way to make things work while keeping the config strongly-typed
type Config struct {
}

type IConfig interface {
	Load() (*Config, error)
	Save(cfg *Config) error
	Remove() error
	Location() (string, error)
}

func NewCfgHandler() IConfig {
	// TODO: Find a way to dynamically return a config handler from the 2 possible options (ie. hostCfgHandler or pluginCfgHandler).
	// aka Strategy pattern

	cfg := &hostCfgHandler{}
	// cfg := &pluginCfgHandler{}

	return cfg
}

func hasCustomLocation() bool {
	cfg := os.Getenv(envName)
	return cfg != ""
}

func defaultDir() (string, error) {
	userCfgDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(userCfgDir, "charmil"), nil
}

// readFile reads the file specified by filePath and returns its contents.
func readFile(filePath string) ([]byte, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return nil, err
	}
	if err != nil {
		return nil, fmt.Errorf("%v: %w", "unable to check if config file exists", err)
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("%v: %w", "unable to read config file", err)
	}

	return data, nil
}

// writeFile writes data to the file specified by filePath.
func writeFile(filePath string, data []byte) error {
	err := ioutil.WriteFile(filePath, data, 0600)
	if err != nil {
		return err
	}

	return nil
}

func getJsonSubkey(data []byte, keyName string) []byte {
	val := gjson.Get(string(data), keyName)

	return []byte(val.String())
}

func setJsonSubkey(data []byte, keyName string, cfg *Config) ([]byte, error) {
	val, err := sjson.Set(string(data), keyName, cfg)
	if err != nil {
		return nil, err
	}

	return []byte(val), nil
}

func deleteJsonSubkey(data []byte, keyName string) ([]byte, error) {
	val, err := sjson.Delete(string(data), keyName)
	if err != nil {
		return nil, err
	}

	return []byte(val), nil
}
