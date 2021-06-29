package rules

import (
	"github.com/aerogear/charmil/validator"
	"github.com/spf13/cobra"
)

// A rule executor which will execute the selected rules
// for now all rules!!
// user API is like this
/*
	validationErrors := validator.executeRules(cmd)
*/

type RuleConfig struct {
	Verbose bool
	Length
	MustExist
}

func (config *RuleConfig) ExecuteRules(cmd *cobra.Command) []validator.ValidationError {
	var errors []validator.ValidationError
	info := validator.StatusLog{TotalTested: 0, TotalErrors: 0, Errors: errors}

	config.initDefaultRules()

	return config.executeHelper(cmd, info)
}

func (config *RuleConfig) executeHelper(cmd *cobra.Command, info validator.StatusLog) []validator.ValidationError {

	return validator.Traverse(
		cmd,
		config.Verbose,
		info,
		config,
		func(cmd *cobra.Command, verbose bool) []validator.ValidationError {
			info.Errors = append(info.Errors, validateLength(cmd, config.Length, config.Verbose)...)
			info.Errors = append(info.Errors, validateMustExist(cmd, config.MustExist, config.Verbose)...)
			return info.Errors
		},
	)
}

func (config *RuleConfig) initDefaultRules() {
	if config.Length.Enable && config.Length.Limits == nil {
		config.Length.Limits = map[string]Limit{
			"Use":     {Min: 2, Max: 20},
			"Short":   {Min: 15, Max: 200},
			"Long":    {Min: 100, Max: 10000},
			"Example": {Min: 100, Max: 10000},
		}
	}

	if config.MustExist.Enable && config.MustExist.Fields == nil {
		config.MustExist.Fields = []string{"Use", "Short", "Long", "Example"}
	}
}
