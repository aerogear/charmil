package rules

import (
	"encoding/json"
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
	Verbose bool `json:"Verbose"`
}

// ValidatorRules consists of all the rules
// present in validator
type ValidatorRules struct {
	Length    `json:"Length"`
	MustExist `json:"MustExist"`
}

// ExecuteRules executes all the rules
// provided by validatorConfig
func (validatorConfig *ValidatorConfig) ExecuteRules(cmd *cobra.Command) []validator.ValidationError {
	var ruleConfig RuleConfig
	return ruleConfig.ExecuteRulesInternal(cmd, validatorConfig)
}

// ValidatorConfigToRuleConfig intializes the default config
// and overrides default with user provided config
func ValidatorConfigToRuleConfig(validatorConfig *ValidatorConfig, ruleConfig *RuleConfig) {
	defaultVerbose := validatorConfig.ValidatorOptions.Verbose

	defaultConfigJson := `{
		"ValidatorOptions": {
			"Verbose": false
		},
		"ValidatorRules": {
			"Length": {
				"Limits": {
					"Use":     {"Min": 2},
					"Short":   {"Min": 15},
					"Long":    {"Min": 50},
					"Example": {"Min": 50}
				}
			},
			"MustExist": {
				"Fields": {"Use": true, "Short": true, "Long": true, "Example": true}
			}
		}
	}`

	// unmarshal defaultConfigJson in configHelper
	var configHelper ValidatorConfig
	if err := json.Unmarshal([]byte(defaultConfigJson), &configHelper); err != nil {
		log.Fatal(err)
	}

	configHelper.Length.Verbose = defaultVerbose
	configHelper.MustExist.Verbose = defaultVerbose

	// Merge user provided config into configHelper
	if err := mergo.Merge(&configHelper, validatorConfig, mergo.WithSliceDeepCopy); err != nil {
		log.Fatal(err)
	}
	validatorConfig = &configHelper

	// append rules to execute
	ruleConfig.Rules = append(ruleConfig.Rules, &validatorConfig.Length)
	ruleConfig.Rules = append(ruleConfig.Rules, &validatorConfig.MustExist)
}
