package crud

import (
	"fmt"
	"os"
	"text/template"

	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/pkg/template/crud"
	"github.com/spf13/cobra"
)

// FlagVariables defines variables that will store the local flag values
type FlagVariables struct {
	// Stores value of the `path` local flag. Default Value: "."
	path string

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
			return generateCrudFiles(flagVars)
		},
	}

	// Adds local flags
	cmd.Flags().StringVar(&flagVars.path, "path", ".", "path where CRUD files need to be generated")
	cmd.Flags().StringVarP(&flagVars.Singular, "singular", "s", "", "name in singular form")
	cmd.Flags().StringVarP(&flagVars.Plural, "plural", "p", "", "name in plural form")

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

// generateCrudFiles generates the CRUD files in the path specified by the `path` flag
func generateCrudFiles(flagVars FlagVariables) error {

	// Creates a folder named `crud` in the path specified using flag
	err := os.Mkdir(flagVars.path+"/crud", 0755)
	if err != nil {
		return err
	}

	// Generates CRUD files in the newly created `crud` folder
	for tmplName, tmplContent := range crud.TmplMap {
		crudFile, err := os.Create(fmt.Sprintf("%s/crud/%s.go", flagVars.path, tmplName))
		if err != nil {
			return err
		}
		defer crudFile.Close()

		crudTemplate := template.Must(template.New(tmplName).Parse(tmplContent))
		err = crudTemplate.Execute(crudFile, flagVars)
		if err != nil {
			return err
		}
	}

	return nil
}
