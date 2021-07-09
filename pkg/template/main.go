package template

func MainTemplate() []byte {
	return []byte(`
package main

import (
	"context"
	"embed"
	"fmt"
	"os"

	"github.com/aerogear/charmil/core/build"
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/localize"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"

	"{{ .PkgName }}/pkg/cmd/root"
)

//go:embed locales
var defaultPath embed.FS

func {{ .CliName }}() (*cobra.Command, *factory.Factory) {
	// init localizer
	localizer, err := localize.New(&localize.Config{
		Files:    defaultPath,
		Language: &language.English,
		Format:   "yaml",
	})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	buildVersion := build.Version
	// init factory
	cmdFactory := factory.Default(localizer)

	// root command
	rootCmd := root.NewRootCommand(cmdFactory, buildVersion)
	rootCmd.InitDefaultHelpCmd()

	return rootCmd, cmdFactory
}

func main() {
	cmd, cmdFactory := abc()

	cmd.Execute()
	build.CheckForUpdate(context.Background(), cmdFactory.Logger, cmdFactory.Localizer)
}
`)
}

func RootTemplate() []byte {
	return []byte(`
package root

import (
	"github.com/aerogear/charmil/core/factory"
	"github.com/spf13/cobra"

	Cliversion "{{ .PkgName }}/pkg/cmd/version"
)

func NewRootCommand(f *factory.Factory, version string) *cobra.Command {

	cmd := &cobra.Command{
		Use:           f.Localizer.LocalizeByID("root.cmd.use"),
		Short:         f.Localizer.LocalizeByID("root.cmd.short"),
		Long:          f.Localizer.LocalizeByID("root.cmd.long"),
		Example:       f.Localizer.LocalizeByID("root.cmd.example"),
		SilenceUsage:  true,
		SilenceErrors: true,
	}

	cmd.Version = version

	cmd.AddCommand(Cliversion.NewVersionCmd(f))

	return cmd
}	
`)
}

func VersionTemplate() []byte {
	return []byte(`
package version

import (
	"fmt"

	"github.com/aerogear/charmil/core/build"
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/core/iostreams"
	"github.com/aerogear/charmil/core/localize"
	"github.com/aerogear/charmil/core/logging"
	"github.com/spf13/cobra"
)

type Options struct {
	IO        *iostreams.IOStreams
	Logger    logging.Logger
	localizer localize.Localizer
}

func NewVersionCmd(f *factory.Factory) *cobra.Command {
	opts := &Options{
		IO:        f.IOStreams,
		Logger:    f.Logger,
		localizer: f.Localizer,
	}

	cmd := &cobra.Command{
		Use:    opts.localizer.LocalizeByID("version.cmd.use"),
		Short:  opts.localizer.LocalizeByID("version.cmd.short"),
		Hidden: true,
		Args:   cobra.NoArgs,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCmd(opts)
		},
	}

	return cmd
}

func runCmd(opts *Options) (err error) {
	fmt.Fprintln(opts.IO.Out, opts.localizer.LocalizeByID("version.cmd.outputText", localize.NewEntry("Version", build.Version)))
	return nil
}

`)
}
