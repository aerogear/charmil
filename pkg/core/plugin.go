package core

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var plugins = make(map[string]Plugin)

func Register(name string, plugin Plugin) {
	if plugin == nil {
		log.Fatal("Undefined plugin")
	}
	if _, isDup := plugins[name]; isDup {
		log.Fatal("Plugin with this name already exists")
	}
	plugins[name] = plugin
}

func GetRootCmd(pluginName string) (*cobra.Command, error) {
	p, ok := plugins[pluginName]
	if !ok {
		return nil, fmt.Errorf("Plugin not found")
	}
	cmd := p.CreateRootCmd()

	return cmd, nil
}
