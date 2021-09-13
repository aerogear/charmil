package init

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/aerogear/charmil/cli/pkg/factory"
	"github.com/aerogear/charmil/core/utils/color"
	"github.com/go-git/go-git/v5"
	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

// promptContent is a struct to hold the prompt content
type promptContent struct {
	errorMsg string
	label    string
}

// TemplateContext contains the context for rendering templates
type TemplateContext struct {
	Owner   string
	Repo    string
	CliName string
}

// InitCommand initializes the starter project
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

			cloneStarter(f)

			f.Logger.Infoln(color.Info("updating starter code with names"))

			// Searches the generated files for default values and replaces it with the user-specified values
			if err := replaceText(templateContext); err != nil {
				f.Logger.Error(err)
				os.Exit(1)
			}

			f.Logger.Infof(color.Success("Your %s CLI has been initialized in this directory.\n"), templateContext.CliName)
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

	_, cloneErr := git.PlainClone(path, false, &git.CloneOptions{
		URL:      "https://github.com/aerogear/charmil-starter",
		Progress: f.IOStreams.Out,
	})
	if cloneErr != nil {
		f.Logger.Error(cloneErr)
		os.Exit(1)
	}
}

// replaceText replaces the default values in files with the user-specified values
func replaceText(templateContext TemplateContext) error {
	replacerList := []string{
		"aerogear/charmil-starter", fmt.Sprintf("%s/%s", templateContext.Owner, templateContext.Repo),
		"placeholdercli", templateContext.CliName,
		"[name of copyright owner]", templateContext.Owner,
	}

	rep := strings.NewReplacer(replacerList...)

	path, err := os.Getwd()
	if err != nil {
		return err
	}

	// rename cli name as given by user
	oldPath := path + "/cmd/placeholdercli"
	newPath := path + "/cmd/" + templateContext.CliName
	if err = os.Rename(oldPath, newPath); err != nil {
		return err
	}

	// remove .git folder
	err = os.RemoveAll(path + "/.git")
	if err != nil {
		return err
	}

	err = filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// Skips directories
			if info.IsDir() || info.Name() == "." {
				return nil
			}

			// Stores contents of the current file
			fileContents, err := ioutil.ReadFile(path)
			if err != nil {
				return err
			}

			// Replaces default values
			updatedFileContents := rep.Replace(string(fileContents))

			// Writes the updated contents to the current file
			err = ioutil.WriteFile(path, []byte(updatedFileContents), 0600)
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
