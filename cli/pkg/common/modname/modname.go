package modname

import (
	"io/ioutil"

	"golang.org/x/mod/modfile"
)

// GetModuleName returns the module name extracted from the `go.mod` file
func GetModuleName() (string, error) {
	// Stores the contents of `go.mod` file as a byte array
	goModBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		return "", err
	}

	// Extracts module name from the passed `go.mod` file contents
	modName := modfile.ModulePath(goModBytes)

	return modName, nil
}
