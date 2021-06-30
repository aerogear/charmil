package rules

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aerogear/charmil/validator"
	"github.com/spf13/cobra"
)

// RuleConfig is the struct that stores
// configuration of rules
type RuleConfig struct {
	Verbose bool
	Length
	MustExist
}

// ExecuteRules executes all the rules
// according to the RuleConfig provided
func (config *RuleConfig) ExecuteRules(cmd *cobra.Command) []validator.ValidationError {
	var errors []validator.ValidationError
	info := validator.StatusLog{TotalTested: 0, TotalErrors: 0, Errors: errors}

	// initialize default rules
	config.initDefaultRules()

	// validate the root command
	config.validate(cmd, &info)

	return config.executeHelper(cmd, &info)
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

	// Length rule
	lenErrs := validateLength(cmd, config)
	info.TotalErrors += len(lenErrs)

	// MustExist rule
	mustExistErrs := validateMustExist(cmd, config)
	info.TotalErrors += len(mustExistErrs)

	info.Errors = append(info.Errors, lenErrs...)
	info.Errors = append(info.Errors, mustExistErrs...)
	info.TotalTested++
}

// initDefaultRules initialize default rules
// and overrides the default rules if RuleConfig is provided by the user
func (config *RuleConfig) initDefaultRules() {

	// default config for rules
	var defaultConfig = &RuleConfig{
		Verbose: false,
		Length: Length{
			Limits: map[string]Limit{
				"Use":     {Min: 2},
				"Short":   {Min: 15},
				"Long":    {Min: 100},
				"Example": {Min: 100},
			},
		},
		MustExist: MustExist{
			Fields: []string{"Use", "Short", "Long", "Example"},
		},
	}

	// Check verbose input from user
	var verbose bool
	if config != nil && config.Verbose {
		verbose = true
	}

	// Set Config to defaultConfig
	*config = *defaultConfig
	config.Verbose = verbose

	// Merge the defaultConfig and Config given by user
	if config.Length.Limits != nil && config.MustExist.Fields != nil {
		out, err := json.Marshal(config)
		if err != nil {
			log.Fatal(err)
		}
		data := []byte(out)

		errr := json.Unmarshal(data, &defaultConfig)
		if errr != nil {
			log.Fatal(errr)
		}
	}

}
