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
	Rules []Rules
}

// ExecuteRulesInternal executes all the rules
// provided by ruleConfig
func ExecuteRulesInternal(cmd *cobra.Command, ruleConfig *RuleConfig, userValidatorConfig *ValidatorConfig) []validator.ValidationError {
	var errors []validator.ValidationError
	info := validator.StatusLog{TotalTested: 0, TotalErrors: 0, Errors: errors}

	// if command's children needs to be ignored
	// only for root
	if skipCommand(userValidatorConfig, cmd.CommandPath()+"*") {
		return info.Errors
	}

	// initialize default rules
	initDefaultRules(userValidatorConfig, ruleConfig)

	// validate the root command
	validate(cmd, &info, ruleConfig, userValidatorConfig)

	return executeHelper(cmd, &info, ruleConfig, userValidatorConfig)
}

func executeHelper(cmd *cobra.Command, info *validator.StatusLog, ruleConfig *RuleConfig, userValidatorConfig *ValidatorConfig) []validator.ValidationError {

	info.Errors = executeRecursive(cmd, info, ruleConfig, userValidatorConfig)

	// prints additional info for the checks
	fmt.Fprintf(os.Stderr, "commands checked: %d\nchecks failed: %d\n", info.TotalTested, info.TotalErrors)

	return info.Errors
}

// executeRecursive recursively traverse over all the subcommands
// and validate using executeRulesChildren function
func executeRecursive(cmd *cobra.Command, info *validator.StatusLog, ruleConfig *RuleConfig, userValidatorConfig *ValidatorConfig) []validator.ValidationError {

	for _, child := range cmd.Commands() {
		// base case
		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}
		// recursive call
		info.Errors = executeRecursive(child, info, ruleConfig, userValidatorConfig)
	}
	info.Errors = executeRulesChildren(cmd, info, ruleConfig, userValidatorConfig)

	return info.Errors
}

// executeRulesChildren execute rules on children of cmd
func executeRulesChildren(cmd *cobra.Command, info *validator.StatusLog, ruleConfig *RuleConfig, userValidatorConfig *ValidatorConfig) []validator.ValidationError {

	// if command's children needs to be ignored
	if skipCommand(userValidatorConfig, cmd.CommandPath()+"*") {
		return info.Errors
	}

	children := cmd.Commands()
	for _, child := range children {

		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}
		validate(child, info, ruleConfig, userValidatorConfig)
	}
	return info.Errors
}

// validate returns validation errors by executing the rules
func validate(cmd *cobra.Command, info *validator.StatusLog, ruleConfig *RuleConfig, userValidatorConfig *ValidatorConfig) {

	// if command needs to be ignored
	if skipCommand(userValidatorConfig, cmd.CommandPath()) {
		return
	}

	// if command's children needs to be ignored
	if skipCommand(userValidatorConfig, cmd.CommandPath()+"*") {
		return
	}

	// traverse all rules and validate
	for _, rule := range ruleConfig.Rules {
		validationErrors := rule.Validate(cmd)
		info.TotalErrors += len(validationErrors)
		info.Errors = append(info.Errors, validationErrors...)
		info.TotalTested++
	}

}

// initDefaultRules initialize default rules
// and overrides the default rules if RuleConfig is provided by the user
func initDefaultRules(validatorConfig *ValidatorConfig, ruleConfig *RuleConfig) {
	ValidatorConfigToRuleConfig(validatorConfig, ruleConfig)
}

// skipCommand checks that command need to be skipped or not
func skipCommand(userValidatorConfig *ValidatorConfig, cmdPath string) bool {
	val, ok := userValidatorConfig.ValidatorOptions.SkipCommands[cmdPath]
	if !ok {
		return ok
	}
	return val
}
