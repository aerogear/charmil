package crud

import (
	"fmt"
	"html/template"
	"io/fs"
	"os"

	"github.com/aerogear/charmil/cli/internal/template/crud"
	"github.com/aerogear/charmil/core/factory"
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
			return generateCrudFiles()
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

// generateCrudFiles generates the CRUD files in the path specified by the `crudPath` flag
func generateCrudFiles() error {
	// Stores path of the directory named `crud` that
	// will be created to store generated CRUD files
	crudDirPath := flagVars.crudPath + "/crud"

	// Creates a directory using value in the `crudDirPath` variable
	err := os.MkdirAll(crudDirPath, 0755)
	if err != nil {
		return err
	}

	// Generates CRUD files in the `crud` directory by looping through the template files
	err = fs.WalkDir(crud.CrudTemplates, ".", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Excludes non-template files from generation
		if info.Name() == "." || info.Name() == "tmpl.go" {
			return nil
		}

		// Stores the current template file contents as a byte array
		buf, err := fs.ReadFile(crud.CrudTemplates, info.Name())
		if err != nil {
			return err
		}

		// Generate CRUD file from the current template
		err = generateFileFromTemplate(info.Name(), crudDirPath, string(buf), flagVars)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// generateFileFromTemplate uses the passed contents and data object of a
// template to generate a new file using the specified file name and output path
func generateFileFromTemplate(name, path, tmplContent string, tmplData interface{}) error {
	// Sets appropriate target path for the locale file
	if name == "crud.en.yaml" && flagVars.localePath != "." {
		// Creates all necessary parent directories from the specified locale path
		err := os.MkdirAll(flagVars.localePath, 0755)
		if err != nil {
			return err
		}

		// Sets the target path
		path = flagVars.localePath
	}

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
