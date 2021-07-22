package crud

import (
	"fmt"
	"os"
	"text/template"

	"github.com/aerogear/charmil/core/factory"
	tmpl "github.com/aerogear/charmil/pkg/template"
	"github.com/spf13/cobra"
)

type FlagVariables struct {
	path string

	Singular string
	Plural   string
}

var (
	flagVars = FlagVariables{}

	tmplMap = map[string]func() []byte{
		"create":   tmpl.CreateCrudTemplate,
		"delete":   tmpl.DeleteCrudTemplate,
		"describe": tmpl.DescribeCrudTemplate,
		"list":     tmpl.ListCrudTemplate,
		"use":      tmpl.UseCrudTemplate,
	}
)

func CrudCommand(f *factory.Factory) *cobra.Command {
	cmd := &cobra.Command{
		Use:           f.Localizer.LocalizeByID("crud.cmd.use"),
		Short:         f.Localizer.LocalizeByID("crud.cmd.short"),
		Long:          f.Localizer.LocalizeByID("crud.cmd.long"),
		Example:       f.Localizer.LocalizeByID("crud.cmd.example"),
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return generateCrudFiles(tmplMap, flagVars)
		},
	}

	cmd.Flags().StringVar(&flagVars.path, "path", ".", "path where CRUD files need to be generated")
	cmd.Flags().StringVarP(&flagVars.Singular, "singular", "s", "", "name in singular form")
	cmd.Flags().StringVarP(&flagVars.Plural, "plural", "p", "", "name in plural form")

	cmd.MarkFlagRequired("singular")
	cmd.MarkFlagRequired("plural")

	return cmd
}

func generateCrudFiles(tmplMap map[string]func() []byte, flagVars FlagVariables) error {

	err := os.Mkdir(flagVars.path+"/crud", 0755)
	if err != nil {
		return err
	}

	for tmplName, tmplFunc := range tmplMap {
		crudFile, err := os.Create(fmt.Sprintf("%s/crud/%s.go", flagVars.path, tmplName))
		if err != nil {
			return err
		}
		defer crudFile.Close()

		crudTemplate := template.Must(template.New(tmplName).Parse(string(tmplFunc())))
		err = crudTemplate.Execute(crudFile, flagVars)
		if err != nil {
			return err
		}
	}

	return nil
}
