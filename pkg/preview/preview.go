package preview

import (
	"log"
	"os"
	"strconv"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

type devPreview struct {
	// Stores all the commands that have been added for dev preview.
	commands []*cobra.Command

	// Holds a `true` value if the `dev_preview` flag has been set.
	isEnabled bool
}

// Creates an instance with false as the default flag value
var dp = &devPreview{isEnabled: false}

// InitFlag initializes a `dev_preview` flag and adds it to the command passed in as argument.
func InitFlag(rootCmd *cobra.Command) {
	// Updates `dp.isEnabled` with the current value present in config file before proceeding
	loadConfig()

	// Adds local flag to the specified command
	rootCmd.Flags().BoolVarP(&dp.isEnabled, "dev_preview", "d", dp.isEnabled, "Enable dev preview commands")
}

// loadConfig fetches env var value present in the config file
// and syncs it with `dp.isEnabled`.
func loadConfig() {
	// Loads the config file
	err := godotenv.Load("./config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Fetches the env var value from config
	boolVal, err := strconv.ParseBool(os.Getenv("CHARMIL_DEV_PREVIEW_ENABLED"))
	if err != nil {
		log.Fatal(err)
	}

	// Updates `dp.isEnabled` with the fetched value
	dp.isEnabled = boolVal
}

// UpdateFlagValue updates the env var value in the config file
// whenever the `dev_preview` flag's value is changed.
func UpdateFlagValue(cmd *cobra.Command, args []string) {
	if cmd.Flags().Changed("dev_preview") {
		env, err := godotenv.Unmarshal("CHARMIL_DEV_PREVIEW_ENABLED=" + strconv.FormatBool(dp.isEnabled))
		if err != nil {
			log.Fatal("Unable to unmarshal")
		}
		err = godotenv.Write(env, "./config.env")
		if err != nil {
			log.Fatal("Unable to write file")
		}

		// Updates the hidden status of every preview command
		updateCmdVisibility()
	}
}

// AddCommands stores the commands passed in as arguments
// and updates their fields.
func AddCommands(cmdList ...*cobra.Command) {
	dp.commands = cmdList
	for _, cmd := range dp.commands {
		updateCmdFields(cmd)
	}
}

// updateCmdFields modifies fields of the command passed in
// as argument with appropriate values to mark it as preview.
func updateCmdFields(cmd *cobra.Command) {
	cmd.Annotations = map[string]string{"channel": "preview"}
	cmd.Short = color.HiCyanString("[Preview] ") + cmd.Short
	cmd.Hidden = !dp.isEnabled
}

// updateCmdVisibility loops through all the stored commands
// and updates their hidden status.
func updateCmdVisibility() {
	for _, cmd := range dp.commands {
		cmd.Hidden = !dp.isEnabled
	}
}
