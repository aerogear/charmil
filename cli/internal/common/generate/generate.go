package generate

import (
	"fmt"
	"html/template"
	"os"
)

// GenerateFileFromTemplate uses the template to generate a
// new file using the specified file name and output path
func GenerateFileFromTemplate(name, path, tmplContent string, placeholderData interface{}) error {
	// Creates a new file using the specified name and path
	f, err := os.Create(fmt.Sprintf("%s/%s", path, name))
	if err != nil {
		return err
	}
	defer f.Close()

	// Adds content to the generated file using the specified template
	tmpl := template.Must(template.New(name).Parse(tmplContent))
	err = tmpl.Execute(f, placeholderData)
	if err != nil {
		return err
	}

	return nil
}
