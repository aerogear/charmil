package crud

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/aerogear/charmil/cli/pkg/common/generate"
	"github.com/aerogear/charmil/cli/pkg/common/modname"

	"github.com/aerogear/charmil/cli/internal/template/crud"
	"github.com/aerogear/charmil/cli/pkg/factory"
	"github.com/spf13/cobra"
)

// TemplateData defines fields that will store all the data used for generating templates
type TemplateData struct {
	// Stores value of the `crudPath` local flag. Default Value: "."
	crudPath string

	// Stores value of the `localePath` local flag. Default Value: "."
	localePath string

	// Stores value of the `singular` local flag
	Singular string

	// Stores value of the `plural` local flag
	Plural string

	// Stores the name of the root module (extracted from go.mod file)
	ModName string
}

// Initializes a zero-valued struct
var tmplData = TemplateData{}

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
			// Extracts the module name from `go.mod` file and stores it
			modName, err := modname.GetModuleName()
			if err != nil {
				return err
			}
			tmplData.ModName = modName

			return generateCrudPackages()
		},
	}

	// Adds local flags
	cmd.Flags().StringVarP(&tmplData.crudPath, f.Localizer.LocalizeByID("crud.flag.crudPath.name"), "c", ".", f.Localizer.LocalizeByID("crud.flag.crudPath.description"))
	cmd.Flags().StringVarP(&tmplData.localePath, f.Localizer.LocalizeByID("crud.flag.localePath.name"), "l", ".", f.Localizer.LocalizeByID("crud.flag.localePath.description"))
	cmd.Flags().StringVarP(&tmplData.Singular, f.Localizer.LocalizeByID("crud.flag.singular.name"), "s", "", f.Localizer.LocalizeByID("crud.flag.singular.description"))
	cmd.Flags().StringVarP(&tmplData.Plural, f.Localizer.LocalizeByID("crud.flag.plural.name"), "p", "", f.Localizer.LocalizeByID("crud.flag.plural.description"))

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

		// Generates the language and the root command files
		if info.Name() == "crud.en.yaml" || info.Name() == "root.tmpl" {
			if err = generateCrudFile(info.Name(), ".", tmplData.crudPath); err != nil {
				return err
			}

			return nil
		}

		if !info.IsDir() || info.Name() == "." {
			return nil
		}

		entries, err := crud.CrudTemplates.ReadDir(path)
		if err != nil {
			return err
		}

		// Stores the path where the current CRUD file will be generated
		targetPath := fmt.Sprintf("%s/%s", tmplData.crudPath, info.Name())

		// Generates CRUD files in separate packages
		for _, entry := range entries {
			// Ensures all parent directories in `targetPath` are created before file generation
			err := os.MkdirAll(targetPath, 0755)
			if err != nil {
				return err
			}

			// Generates CRUD files in the corresponding packages
			err = generateCrudFile(entry.Name(), path, targetPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

// generateCrudFile takes the target file name, target path and the path
// of the template as arguments and generates a CRUD file using it.
func generateCrudFile(fileName, currentPath, targetPath string) error {
	// Stores the current template file contents as a byte array
	buf, err := crud.CrudTemplates.ReadFile(filepath.Join(currentPath, fileName))
	if err != nil {
		return err
	}

	switch fileName {
	case "crud.en.yaml":
		if tmplData.localePath != "." {
			targetPath = tmplData.localePath

			// Ensures all parent directories in `targetPath` are created before file generation
			if err = os.MkdirAll(targetPath, 0755); err != nil {
				return err
			}
		}

		fileName = tmplData.Singular + "." + fileName

	case "root.tmpl":
		fileName = tmplData.Singular + ".go"
	default:
		fileName = fileName[:len(fileName)-5] + ".go"
	}

	// Generate CRUD file from the current template
	err = generate.GenerateFileFromTemplate(fileName, targetPath, string(buf), tmplData)
	if err != nil {
		return err
	}

	return nil
}
