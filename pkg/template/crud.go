package template

func CreateCrudTemplate() []byte {
	return []byte(`// Create file for {{.Singular}} instances

package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type createOptions struct {
	// Add your option fields here
}

// NewCreateCommand creates a new command for creating instances.
func NewCreateCommand(f *factory.Factory) *cobra.Command {
	opts := &createOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("crud.cmd.create.use"),
		Short:   f.Localizer.LocalizeByID("crud.cmd.create.shortDescription"),
		Long:    f.Localizer.LocalizeByID("crud.cmd.create.longDescription"),
		Example: f.Localizer.LocalizeByID("crud.cmd.create.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runCreate(opts, f)
		},
	}

	return cmd
}

func runCreate(opts *createOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
`)
}

func DeleteCrudTemplate() []byte {
	return []byte(`// Delete file for {{.Singular}} instances
	
package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type deleteOptions struct {
	// Add your option fields here

}

// NewDeleteCommand creates a new command for deleting instances.
func NewDeleteCommand(f *factory.Factory) *cobra.Command {
	opts := &deleteOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("crud.cmd.delete.use"),
		Short:   f.Localizer.LocalizeByID("crud.cmd.delete.shortDescription"),
		Long:    f.Localizer.LocalizeByID("crud.cmd.delete.longDescription"),
		Example: f.Localizer.LocalizeByID("crud.cmd.delete.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDelete(opts, f)
		},
	}

	return cmd
}

func runDelete(opts *deleteOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
`)
}

func DescribeCrudTemplate() []byte {
	return []byte(`// Describe file for {{.Singular}} instances

package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type describeOptions struct {
	// Add your option fields here

}

// NewDescribeCommand creates a new command for describing instances.
func NewDescribeCommand(f *factory.Factory) *cobra.Command {
	opts := &describeOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("crud.cmd.describe.use"),
		Short:   f.Localizer.LocalizeByID("crud.cmd.describe.shortDescription"),
		Long:    f.Localizer.LocalizeByID("crud.cmd.describe.longDescription"),
		Example: f.Localizer.LocalizeByID("crud.cmd.describe.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runDescribe(opts, f)
		},
	}

	return cmd
}

func runDescribe(opts *describeOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
`)
}

func ListCrudTemplate() []byte {
	return []byte(`// List file for {{.Singular}} instances
	
package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type listOptions struct {
	// Add your option fields here

}

// NewListCommand creates a new command for listing instances.
func NewListCommand(f *factory.Factory) *cobra.Command {
	opts := &listOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("crud.cmd.list.use"),
		Short:   f.Localizer.LocalizeByID("crud.cmd.list.shortDescription"),
		Long:    f.Localizer.LocalizeByID("crud.cmd.list.longDescription"),
		Example: f.Localizer.LocalizeByID("crud.cmd.list.example"),
		Args:    cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runList(opts, f)
		},
	}

	return cmd
}

func runList(opts *listOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
`)
}

func UseCrudTemplate() []byte {
	return []byte(`// Use file for {{.Singular}} instances
	
package crud

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"
)

type useOptions struct {
	// Add your option fields here

}

// NewUseCommand creates a new command for using instances.
func NewUseCommand(f *factory.Factory) *cobra.Command {
	opts := &useOptions{}

	cmd := &cobra.Command{
		Use:     f.Localizer.LocalizeByID("crud.cmd.use.use"),
		Short:   f.Localizer.LocalizeByID("crud.cmd.use.shortDescription"),
		Long:    f.Localizer.LocalizeByID("crud.cmd.use.longDescription"),
		Example: f.Localizer.LocalizeByID("crud.cmd.use.example"),
		Args:    cobra.RangeArgs(0, 1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runUse(opts, f)
		},
	}

	return cmd
}

func runUse(opts *useOptions, f *factory.Factory) error {
	// Add your implementation here

	return nil
}
`)
}
