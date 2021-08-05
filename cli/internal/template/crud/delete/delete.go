// Delete file for {{.Singular}} instances

package delete

import (
	"errors"

	"github.com/aerogear/charmil/cli/internal/factory"
	"github.com/spf13/cobra"
)

type deleteOptions struct {
	// Stores value of the `name` argument
	name string

	// Stores value of the `id` flag
	id string

	// You can add more fields here according to your requirements
}

// GetDeleteCommand returns a new command for deleting {{.Singular}} instances.
func GetDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &deleteOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("{{.Singular}}.cmd.delete.use"),
		Short:   f.Localizer.LocalizeByID("{{.Singular}}.cmd.delete.shortDescription"),
		Long:    f.Localizer.LocalizeByID("{{.Singular}}.cmd.delete.longDescription"),
		Example: f.Localizer.LocalizeByID("{{.Singular}}.cmd.delete.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {

			if len(args) > 0 {
				opts.name = args[0]
			}

			if err := performValidation(opts, f); err != nil {
				return err
			}

			return runDelete(opts, f)
		},
	}

	// Adds local flags
	cmd.Flags().StringVar(&opts.id, "id", "", f.Localizer.LocalizeByID("{{.Singular}}.common.flag.id"))

	return cmd
}

// performValidation validates the arguments and flag values of the Delete command
func performValidation(opts *deleteOptions, f *factory.Factory) error {

	// Ensures that at least one of the name and id values are specified
	if opts.name == "" && opts.id == "" {
		return errors.New(f.Localizer.LocalizeByID("{{.Singular}}.error.idOrNameRequired"))
	}

	// Returns an error when both name and id values are specified together
	if opts.name != "" && opts.id != "" {
		return errors.New(f.Localizer.LocalizeByID("{{.Singular}}.error.idAndNameCannotBeUsed"))
	}

	return nil
}
