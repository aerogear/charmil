package rules

import (
	"fmt"
	"github.com/aerogear/charmil/validator"
	"github.com/spf13/cobra"
	"os"
)

type Rule interface {
	Validate(command *cobra.Command, ruleConfig *RuleConfig) []validator.ValidationError
}

// RuleConfig is the struct that stores
// configuration of rules
type RuleConfig struct {
	Verbose bool
	Length
	MustExist
}

func NewRuleConfig() RuleConfig {
	return RuleConfig{
		Verbose: false,
		Length: Length{
			Limits: map[string]Limit{
				"Use":     {Min: 2},
				"Short":   {Min: 15},
				"Long":    {Min: 50},
				"Example": {Min: 50},
			},
		},
		MustExist: MustExist{
			Fields: []string{"Use", "Short", "Long", "Example"},
		},
	}
}

// ExecuteRules executes all the rules on the command given
// and all child commands
func (ruleConfig *RuleConfig) ExecuteRules(cmd *cobra.Command) []validator.ValidationError {
	var errors []validator.ValidationError
	info := validator.StatusLog{TotalTested: 0, TotalErrors: 0, Errors: errors}

	// validate the root command
	ruleConfig.validate(cmd, &info)

	// validates all children of root
	info.Errors = ruleConfig.executeRecursive(cmd, &info)

	// prints additional info for the checks
	fmt.Fprintf(os.Stderr, "commands checked: %d\nchecks failed: %d\n", info.TotalTested, info.TotalErrors)

	return info.Errors
}

// executeRecursive recursively traverse over all the subcommands
// and validate using executeRulesChildren function
func (ruleConfig *RuleConfig) executeRecursive(cmd *cobra.Command, info *validator.StatusLog) []validator.ValidationError {

	for _, child := range cmd.Commands() {
		// base case
		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}
		// recursive call
		info.Errors = ruleConfig.executeRecursive(child, info)
	}
	info.Errors = ruleConfig.executeRulesChildren(cmd, info)

	return info.Errors
}

// executeRulesChildren execute rules on children of cmd
func (ruleConfig *RuleConfig) executeRulesChildren(cmd *cobra.Command, info *validator.StatusLog) []validator.ValidationError {
	children := cmd.Commands()
	for _, child := range children {

		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}
		ruleConfig.validate(child, info)
	}
	return info.Errors
}

// validate returns validation errors by executing the rules
func (ruleConfig *RuleConfig) validate(command *cobra.Command, info *validator.StatusLog) {

	// could clean this up by having a generic array
	// of rules in the rule config
	var results = [][]validator.ValidationError {
		ruleConfig.Length.Validate(command, ruleConfig),
		ruleConfig.MustExist.Validate(command, ruleConfig),
	}

	for _, result := range results {
		info.TotalErrors += len(result)
		info.Errors = append(info.Errors, result...)
	}
	info.TotalTested++
}
