package crud

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/aerogear/charmil/cli/internal/factory"
	"github.com/aerogear/charmil/cli/internal/template/crud"
	"github.com/spf13/cobra"
)

// FlagVariables defines variables that will store the local flag values
type FlagVariables struct {
	// Stores value of the `crudPath` local flag. Default Value: "."
	crudPath string

	// Stores value of the `localePath` local flag. Default Value: "."
	localePath string

	// Stores value of the `singular` local flag
	Singular string

	// Stores value of the `plural` local flag
	Plural string
}

// Initializes a zero-valued struct
var flagVars = FlagVariables{}

// CrudCommand returns the crud command. This will be
// added to the charmil CLI as a sub-command
func CrudCommand(f *factory.Factory) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:           f.Localizer.LocalizeByID("crud.cmd.use"),
		Short:         f.Localizer.LocalizeByID("crud.cmd.short"),
		Long:          f.Localizer.LocalizeByID("crud.cmd.long"),
		Example:       f.Localizer.LocalizeByID("crud.cmd.example"),
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return generateCrudPackages()
		},
	}

	// Adds local flags
	cmd.Flags().StringVarP(&flagVars.crudPath, f.Localizer.LocalizeByID("crud.flag.crudPath.name"), "c", ".", f.Localizer.LocalizeByID("crud.flag.crudPath.description"))
	cmd.Flags().StringVarP(&flagVars.localePath, f.Localizer.LocalizeByID("crud.flag.localePath.name"), "l", ".", f.Localizer.LocalizeByID("crud.flag.localePath.description"))
	cmd.Flags().StringVarP(&flagVars.Singular, f.Localizer.LocalizeByID("crud.flag.singular.name"), "s", "", f.Localizer.LocalizeByID("crud.flag.singular.description"))
	cmd.Flags().StringVarP(&flagVars.Plural, f.Localizer.LocalizeByID("crud.flag.plural.name"), "p", "", f.Localizer.LocalizeByID("crud.flag.plural.description"))

	// Marks the `singular` flag as required.
	// This causes the crud command to report an
	// error if invoked without the `singular` flag.
	err := cmd.MarkFlagRequired("singular")
	if err != nil {
		return nil, err
	}

	// Marks the `plural` flag as required.
	// This causes the crud command to report an
	// error if invoked without the `plural` flag.
	err = cmd.MarkFlagRequired("plural")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// generateCrudPackages generates the CRUD packages in the path specified by the `crudPath` flag
func generateCrudPackages() error {

	// Generates CRUD packages in the specified path by looping through the template files
	err := fs.WalkDir(crud.CrudTemplates, ".", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Generates the language file
		if info.Name() == "crud.en.yaml" {
			// Sets appropriate target path for the locale file
			if flagVars.localePath == "." {
				flagVars.localePath = flagVars.crudPath
			}

			err = generateCrudFile(info.Name(), ".", flagVars.localePath)
			if err != nil {
				return err
			}

			return nil
		}

		// Generates the root CRUD command file from template
		if info.Name() == "root.go" {
			err = generateFileFromTemplate(flagVars.Singular+".go", flagVars.crudPath, string(crud.RootTemplate), flagVars)
			if err != nil {
				return err
			}
		}

		if !info.IsDir() || info.Name() == "." {
			return nil
		}

		entries, err := crud.CrudTemplates.ReadDir(path)
		if err != nil {
			return err
		}

		// Stores the path where the CRUD file will be generated
		targetPath := fmt.Sprintf("%s/%s", flagVars.crudPath, info.Name())

		// Generates CRUD files in separate packages
		for _, entry := range entries {
			err = generateCrudFile(entry.Name(), path, targetPath)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// generateCrudFile takes the target file name, target path and the path
// of the template as arguments and generates a CRUD file using it.
func generateCrudFile(fileName, currentPath, targetPath string) error {

	// Ensures all parent directories in `targetPath` are created before file generation
	err := os.MkdirAll(targetPath, 0755)
	if err != nil {
		return err
	}

	// Stores the current template file contents as a byte array
	buf, err := crud.CrudTemplates.ReadFile(filepath.Join(currentPath, fileName))
	if err != nil {
		return err
	}

	// Generate CRUD file from the current template
	err = generateFileFromTemplate(fileName, targetPath, string(buf), flagVars)
	if err != nil {
		return err
	}

	return nil
}

// generateFileFromTemplate uses the template to generate a
// new file using the specified file name and output path
func generateFileFromTemplate(name, path, tmplContent string, tmplData interface{}) error {

	// Creates a new file using the specified name and path
	f, err := os.Create(fmt.Sprintf("%s/%s", path, name))
	if err != nil {
		return err
	}
	defer f.Close()

	// Adds content to the generated file using the specified template
	tmpl := template.Must(template.New(name).Parse(tmplContent))
	err = tmpl.Execute(f, tmplData)
	if err != nil {
		return err
	}

	return nil
}
