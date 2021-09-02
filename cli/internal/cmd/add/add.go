package add

import (
	"fmt"
	"html/template"
	"io/fs"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/aerogear/charmil/cli/internal/template/add"
	"github.com/aerogear/charmil/cli/pkg/common/modname"
	"github.com/aerogear/charmil/cli/pkg/factory"
	"github.com/aerogear/charmil/core/utils/color"
	"github.com/spf13/cobra"
)

// TemplateData defines fields that will store all the data used for generating templates
type TemplateData struct {
	// Stores value of the `CmdPath` local flag. Default Value: "."
	CmdPath string

	// Stores value of the `CmdName` local flag
	CmdName string

	// Stores the name of the root module (extracted from go.mod file)
	ModName string
}

// Initializes a zero-valued struct
var tmplData = TemplateData{}

func AddCommand(f *factory.Factory) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:           f.Localizer.LocalizeByID("add.cmd.use"),
		Short:         f.Localizer.LocalizeByID("add.cmd.short"),
		Long:          f.Localizer.LocalizeByID("add.cmd.long"),
		Example:       f.Localizer.LocalizeByID("add.cmd.example"),
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Extracts the module name from `go.mod` file and stores it
			modName, err := modname.GetModuleName()
			if err != nil {
				return err
			}
			tmplData.ModName = modName

			if err := generateCommand(); err != nil {
				return err
			}

			f.Logger.Infof(color.Success("%s command has been created in %s directory\n"), tmplData.CmdName, tmplData.CmdPath+"/"+tmplData.CmdName)

			return nil

		},
	}

	// Adds local flags
	cmd.Flags().StringVarP(&tmplData.CmdPath, f.Localizer.LocalizeByID("add.flag.cmdPath.name"), "c", ".", f.Localizer.LocalizeByID("add.flag.cmdPath.description"))
	cmd.Flags().StringVarP(&tmplData.CmdName, f.Localizer.LocalizeByID("add.flag.cmdName.name"), "s", "", f.Localizer.LocalizeByID("add.flag.cmdName.description"))

	// Marks the `cmdName` flag as required.
	// This causes the add command to report an
	// error if invoked without the `cmdName` flag.
	err := cmd.MarkFlagRequired("cmdName")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

// generateCommand function generates a command with it's locales file
func generateCommand() error {

	// create a directory with command name
	if mkdirerr := os.MkdirAll(path.Join(tmplData.CmdPath, tmplData.CmdName), 0755); mkdirerr != nil {
		return mkdirerr
	}

	// walk through add templates folder
	err := fs.WalkDir(add.AddTemplates, ".", func(p string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// creating templates from add templates folder
		// to the cmdPath directory(provided by user)
		var buf []byte
		if info.Name() == "cmdname.en.yaml" || info.Name() == "cmdname.tmpl" {
			buf, err = add.AddTemplates.ReadFile(info.Name())
			if err != nil {
				return err
			}
			var ext string

			if info.Name() == "cmdname.en.yaml" {
				ext = ".en.yaml"
			} else {
				ext = ".go"
			}

			err = ioutil.WriteFile(path.Join(tmplData.CmdPath, tmplData.CmdName, tmplData.CmdName+ext), buf, 0600)
			if err != nil {
				fmt.Printf("Unable to write file: %v", err)
			}
		}

		// apply templates according to tmplData
		err = applyTemplates()
		return err

	})

	if err != nil {
		return fmt.Errorf("failed to walk directory: %w", err)
	}

	return nil
}

// applyTemplates parses the files and apply templates
func applyTemplates() error {
	err := filepath.Walk(path.Join(tmplData.CmdPath, tmplData.CmdName),
		func(path string, info os.FileInfo, err error) error {

			if err != nil {
				return err
			}

			fi, err := os.Stat(path)
			if err != nil {
				return fmt.Errorf("failed to read file info: %w", err)
			}

			if fi.IsDir() {
				return nil
			}

			tmpl, tmplErr := template.ParseFiles(path)
			if tmplErr != nil {
				return fmt.Errorf("failed to parse template: %w", err)
			}

			f, PathErr := os.Create(path)
			if PathErr != nil {
				return fmt.Errorf("failed to create file: %w", err)
			}

			// apply templateContext to the folder
			if err := tmpl.Execute(f, tmplData); err != nil {
				return fmt.Errorf("failed to execute template: %w", err)
			}

			return nil
		})

	return err
}
