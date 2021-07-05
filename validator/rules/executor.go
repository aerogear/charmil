package rules

import (
	"fmt"
	"os"

	"github.com/aerogear/charmil/validator"
	"github.com/spf13/cobra"
)

// Rules is an interface that is implemented
// by every rule defined in rules package
type Rules interface {
	Validate(cmd *cobra.Command) []validator.ValidationError
}

// RuleConfig is the struct that stores
// configuration of rules
type RuleConfig struct {
	Verbose bool
	Rules   []Rules
}

// ExecuteRulesInternal executes all the rules
// provided by ruleConfig
func (ruleConfig *RuleConfig) ExecuteRulesInternal(cmd *cobra.Command, userValidatorConfig *ValidatorConfig) []validator.ValidationError {
	var errors []validator.ValidationError
	info := validator.StatusLog{TotalTested: 0, TotalErrors: 0, Errors: errors}

	// initialize default rules
	ruleConfig.initDefaultRules(userValidatorConfig)

	// validate the root command
	ruleConfig.validate(cmd, &info)

	return ruleConfig.executeHelper(cmd, &info)
}

func (config *RuleConfig) executeHelper(cmd *cobra.Command, info *validator.StatusLog) []validator.ValidationError {
	info.Errors = config.executeRecursive(cmd, info)

	// prints additional info for the checks
	fmt.Fprintf(os.Stderr, "commands checked: %d\nchecks failed: %d\n", info.TotalTested, info.TotalErrors)

	return info.Errors
}

// executeRecursive recursively traverse over all the subcommands
// and validate using executeRulesChildren function
func (config *RuleConfig) executeRecursive(cmd *cobra.Command, info *validator.StatusLog) []validator.ValidationError {
	for _, child := range cmd.Commands() {
		// base case
		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}
		// recursive call
		info.Errors = config.executeRecursive(child, info)
	}
	info.Errors = config.executeRulesChildren(cmd, info)

	return info.Errors
}

// executeRulesChildren execute rules on children of cmd
func (config *RuleConfig) executeRulesChildren(cmd *cobra.Command, info *validator.StatusLog) []validator.ValidationError {
	children := cmd.Commands()
	for _, child := range children {

		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}
		config.validate(child, info)
	}
	return info.Errors
}

// validate returns validation errors by executing the rules
func (config *RuleConfig) validate(cmd *cobra.Command, info *validator.StatusLog) {

	// traverse all rules and validate
	for _, rule := range config.Rules {
		validationErrors := rule.Validate(cmd)
		info.TotalErrors += len(validationErrors)
		info.Errors = append(info.Errors, validationErrors...)
		info.TotalTested++
	}

}

// initDefaultRules initialize default rules
// and overrides the default rules if RuleConfig is provided by the user
func (config *RuleConfig) initDefaultRules(testCfg *ValidatorConfig) {
	ValidatorConfigToRuleConfig(testCfg, config)
}
