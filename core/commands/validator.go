package commands

import (
	factory2 "github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/validator/rules"
	"github.com/spf13/cobra"
)

func ValidateCommand() (*cobra.Command, error) {

	cmd := &cobra.Command{
		Use:     "validate",
		Short:   "tests rules on added commands",
		Long:    "executes all rules defined on each command added",
		SilenceUsage: false,
		RunE: func(cmd *cobra.Command, args []string) error {

			var ruleConfig rules.RuleConfig
			validationErrors := ruleConfig.ExecuteRules(cmd.Root())

			logger := factory2.Default(nil).Logger
			for _, validationError := range validationErrors {
				if validationError.Err != nil {
					logger.Errorf("%v failed rule %v with error: %v\n", validationError.Cmd.Use,
						validationError.Rule, validationError.Name)
				}
			}
			return nil
		},
	}

	return cmd, nil
}
