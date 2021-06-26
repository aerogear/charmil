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
	commands  []*cobra.Command
	isEnabled bool
}

var dp = &devPreview{isEnabled: false}

func InitFlag(rootCmd *cobra.Command) {
	loadConfig()
	rootCmd.Flags().BoolVarP(&dp.isEnabled, "dev_preview", "d", dp.isEnabled, "Enable dev preview commands")
}

func loadConfig() {
	err := godotenv.Load("./config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	boolVal, err := strconv.ParseBool(os.Getenv("CHARMIL_DEV_PREVIEW_ENABLED"))
	if err != nil {
		log.Fatal(err)
	}

	dp.isEnabled = boolVal
}

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

		updateCmdVisibility()
	}
}

func AddCommands(cmdList ...*cobra.Command) {
	dp.commands = cmdList
	for _, cmd := range dp.commands {
		updateCmdFields(cmd)
	}
}

func updateCmdFields(cmd *cobra.Command) {
	cmd.Annotations = map[string]string{"channel": "preview"}
	cmd.Short = color.HiCyanString("[Preview] ") + cmd.Short
	cmd.Hidden = !dp.isEnabled
}

func updateCmdVisibility() {
	for _, cmd := range dp.commands {
		cmd.Hidden = !dp.isEnabled
	}
}
