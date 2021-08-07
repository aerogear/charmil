package add

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/aerogear/charmil/cli/internal/common/modname"
	"github.com/aerogear/charmil/cli/internal/factory"
	"github.com/aerogear/charmil/cli/internal/template/add"
	"github.com/spf13/cobra"
)

// TemplateData defines fields that will store all the data used for generating templates
type TemplateData struct {
	// Stores value of the `cmdPath` local flag. Default Value: "."
	cmdPath string

	// Stores value of the `localePath` local flag. Default Value: "."
	localePath string

	// Stores value of the `CmdName` local flag
	cmdName string

	// Stores the name of the root module (extracted from go.mod file)
	modName string
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
			tmplData.modName = modName

			return generateCommand()
		},
	}

	// Adds local flags
	cmd.Flags().StringVarP(&tmplData.cmdPath, f.Localizer.LocalizeByID("add.flag.cmdPath.name"), "c", ".", f.Localizer.LocalizeByID("add.flag.cmdPath.description"))
	cmd.Flags().StringVarP(&tmplData.localePath, f.Localizer.LocalizeByID("add.flag.localePath.name"), "l", ".", f.Localizer.LocalizeByID("add.flag.localePath.description"))
	cmd.Flags().StringVarP(&tmplData.cmdName, f.Localizer.LocalizeByID("add.flag.cmdName.name"), "s", "", f.Localizer.LocalizeByID("add.flag.cmdName.description"))

	// Marks the `cmdName` flag as required.
	// This causes the add command to report an
	// error if invoked without the `cmdName` flag.
	err := cmd.MarkFlagRequired("cmdName")
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func generateCommand() error {

	err := fs.WalkDir(add.AddTemplates, ".", func(path string, info fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() || info.Name() == "." {
			return nil
		}

		entries, err := add.AddTemplates.ReadDir(path)
		if err != nil {
			return err
		}

		// Stores the path where the current CRUD file will be generated
		targetPath := fmt.Sprintf("%s/%s", tmplData.cmdPath, info.Name())

		// Generates CRUD files in separate packages
		for _, entry := range entries {
			// Ensures all parent directories in `targetPath` are created before file generation
			err := os.MkdirAll(targetPath, 0755)
			if err != nil {
				return err
			}

			// Generates add files in the corresponding packages
			fmt.Println(entry)
		}

		return nil

	})

	return err
}
