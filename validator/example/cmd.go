package main

import (
	"log"

	"github.com/spf13/cobra"
)

type command struct {
	name    *cobra.Command
	use     string
	short   string
	long    string
	example string
}

func NewCommand() *cobra.Command {

	cmd := &cobra.Command{
		Use:     "cmd0",
		Short:   "This is the short",
		Long:    `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis malesuada varius lacus, sit amet dictum risus convallis nec. Quisque suscipit at neque in blandit. Proin a accumsan ante. Aenean cursus suscipit sem. Nunc sollicitudin, ante et vehicula pharetra, mauris elit porta felis, et ultricies nulla justo eleifend justo. Proin sit amet.`,
		Example: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis malesuada varius lacus, sit amet dictum risus convallis nec. Quisque suscipit at neque in blandit. Proin a accumsan ante. Aenean cursus suscipit sem. Nunc sollicitudin, ante et vehicula pharetra, mauris elit porta felis, et ultricies nulla justo eleifend justo. Proin sit amet.",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}

	var commands []command = []command{
		{
			use:     "subcmd01",
			short:   "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis malesuada varius lacus, sit amet dictum risus convallis nec. Quisque suscipit at neque in blandit. Proin a accumsan ante. Aenean cursus suscipit sem. Nunc sollicitudin, ante et vehicula pharetra, mauris elit porta felis, et ultricies nulla justo eleifend justo. Proin sit amet.",
			long:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis malesuada varius lacus, sit amet dictum risus convallis nec. Quisque suscipit at neque in blandit. Proin a accumsan ante. Aenean cursus suscipit sem. Nunc sollicitudin, ante et vehicula pharetra, mauris elit porta felis, et ultricies nulla justo eleifend justo. Proin sit amet.",
			example: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis malesuada varius lacus, sit amet dictum risus convallis nec. Quisque suscipit at neque in blandit. Proin a accumsan ante. Aenean cursus suscipit sem. Nunc sollicitudin, ante et vehicula pharetra, mauris elit porta felis, et ultricies nulla justo eleifend justo. Proin sit amet.",
		},
		{
			use:     "subcmd12",
			short:   "This is the description short12",
			long:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis malesuada varius lacus, sit amet dictum risus convallis nec. Quisque suscipit at neque in blandit. Proin a accumsan ante. Aenean cursus suscipit sem. Nunc sollicitudin, ante et vehicula pharetra, mauris elit porta felis, et ultricies nulla justo eleifend justo. Proin sit amet.",
			example: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis malesuada varius lacus, sit amet dictum risus convallis nec. Quisque suscipit at neque in blandit. Proin a accumsan ante. Aenean cursus suscipit sem. Nunc sollicitudin, ante et vehicula pharetra, mauris elit porta felis, et ultricies nulla justo eleifend justo. Proin sit amet.",
		},
		{
			use:     "subcmd03",
			short:   "This is the description short03",
			long:    "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis malesuada varius lacus, sit amet dictum risus convallis nec. Quisque suscipit at neque in blandit. Proin a accumsan ante. Aenean cursus suscipit sem. Nunc sollicitudin, ante et vehicula pharetra, mauris elit porta felis, et ultricies nulla justo eleifend justo. Proin sit amet.",
			example: "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Duis malesuada varius lacus, sit amet dictum risus convallis nec. Quisque suscipit at neque in blandit. Proin a accumsan ante. Aenean cursus suscipit sem. Nunc sollicitudin, ante et vehicula pharetra, mauris elit porta felis, et ultricies nulla justo eleifend justo. Proin sit amet.",
		},
	}

	for _, cm := range commands {
		cm.name = &cobra.Command{
			Use:     cm.use,
			Short:   cm.short,
			Long:    cm.long,
			Example: cm.example,
			Run:     func(cmd *cobra.Command, args []string) {},
		}
		cmd.AddCommand(cm.name)
	}

	return cmd
}

func main() {
	cmd := NewCommand()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
