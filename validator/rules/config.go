package rules

import (
	"log"

	"github.com/aerogear/charmil/validator"
	"github.com/imdario/mergo"
	"github.com/spf13/cobra"
)

// ValidatorConfig is provided to user for overriding default rules
type ValidatorConfig struct {
	ValidatorOptions `json:"ValidatorOptions"`
	ValidatorRules   `json:"ValidatorRules"`
}

// ValidatorOptions provide additional configurations
// to the rules
type ValidatorOptions struct {
	Verbose      bool            `json:"Verbose"`
	SkipChildren map[string]bool `json:"SkipChildren"`
	SkipCommands map[string]bool `json:"SkipCommands"`
}

// ValidatorRules consists of all the rules
// present in validator
type ValidatorRules struct {
	Length     `json:"Length"`
	MustExist  `json:"MustExist"`
	UseMatches `json:"UseMatches"`
}

// ExecuteRules executes all the rules
// provided by validatorConfig
func ExecuteRules(cmd *cobra.Command, validatorConfig *ValidatorConfig) []validator.ValidationError {
	var ruleConfig RuleConfig
	return ExecuteRulesInternal(cmd, &ruleConfig, validatorConfig)
}

// ValidatorConfigToRuleConfig intializes the default config
// and overrides default with user provided config
func ValidatorConfigToRuleConfig(validatorConfig *ValidatorConfig, ruleConfig *RuleConfig) {
	defaultVerbose := validatorConfig.ValidatorOptions.Verbose

	// unmarshal defaultConfigJson in configHelper
	var configHelper ValidatorConfig = ValidatorConfig{
		ValidatorOptions: ValidatorOptions{
			Verbose: defaultVerbose,
		},
		ValidatorRules: ValidatorRules{
			Length: Length{
				RuleOptions: validator.RuleOptions{
					Verbose: defaultVerbose,
				},
				Limits: map[string]Limit{
					"Use":     {Min: 2},
					"Short":   {Min: 15},
					"Long":    {Min: 50},
					"Example": {Min: 50},
				},
			},
			MustExist: MustExist{
				RuleOptions: validator.RuleOptions{
					Verbose: defaultVerbose,
				},
				Fields: map[string]bool{"Use": true, "Short": true, "Long": true, "Example": true},
			},
		},
	}

	// Merge user provided config into configHelper
	if err := mergo.Merge(&configHelper, validatorConfig, mergo.WithSliceDeepCopy); err != nil {
		log.Fatal(err)
	}
	validatorConfig = &configHelper

	// append rules to execute
	ruleConfig.Rules = append(
		ruleConfig.Rules,
		&validatorConfig.Length,
		&validatorConfig.MustExist,
		&validatorConfig.UseMatches,
	)
}
