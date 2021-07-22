package init

import (
	"errors"
	"fmt"
	"os"

	"github.com/aerogear/charmil/core/factory"
	"github.com/go-git/go-git/v5"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// promptContent is a struct to hold the prompt content
type promptContent struct {
	errorMsg string
	label    string
}

type TemplateContext struct {
	Owner   string
	Repo    string
	CliName string
}

func InitCommand(f *factory.Factory) *cobra.Command {

	cmd := &cobra.Command{
		Use:           f.Localizer.LocalizeByID("init.cmd.use"),
		Short:         f.Localizer.LocalizeByID("init.cmd.short"),
		Long:          f.Localizer.LocalizeByID("init.cmd.long"),
		Example:       f.Localizer.LocalizeByID("init.cmd.example"),
		SilenceErrors: true,
		Run: func(cmd *cobra.Command, args []string) {
			owner := promptGetInput(promptContent{
				label:    "GitHub Organization or Username",
				errorMsg: "Please Provide a username",
			})
			repo := promptGetInput(promptContent{
				label:    "GitHub Repo Name",
				errorMsg: "Please Provide a repo name",
			})
			cli_name := promptGetInput(promptContent{
				label:    "CLI Name",
				errorMsg: "Please Provide a cli name",
			})

			templateContext := TemplateContext{
				Owner:   owner,
				Repo:    repo,
				CliName: cli_name,
			}

			f.Logger.Info(templateContext)

			cloneStarter(f)
		},
	}

	return cmd
}

// promptGetInput returns a string got by prompting the user
func promptGetInput(pc promptContent) string {

	// validate function for validating prompts
	validate := func(input string) error {
		if len(input) == 0 {
			return errors.New(pc.errorMsg)
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    pc.label,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	return result
}

// clone a git repository to a local path using go-git
func cloneStarter(f *factory.Factory) {
	path, pathErr := os.Getwd()
	if pathErr != nil {
		f.Logger.Error(pathErr)
		os.Exit(1)
	}

	_, cloneErr := git.PlainClone(path+"/cloned", false, &git.CloneOptions{
		URL:      "https://github.com/ankithans/charmil-starter-template",
		Progress: f.IOStreams.Out,
	})
	if cloneErr != nil {
		f.Logger.Error(cloneErr)
		os.Exit(1)
	}
}
