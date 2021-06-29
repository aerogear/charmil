package rules

import (
	"encoding/json"
	"log"

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

	return config.executeHelper(cmd, info)
}

func (config *RuleConfig) executeHelper(cmd *cobra.Command, info validator.StatusLog) []validator.ValidationError {

	return validator.Traverse(
		cmd,
		config.Verbose,
		info,
		func(cmd *cobra.Command, verbose bool) []validator.ValidationError {
			info.Errors = append(info.Errors, validateLength(cmd, config.Length, config.Verbose)...)
			info.Errors = append(info.Errors, validateMustExist(cmd, config.MustExist, config.Verbose)...)
			return info.Errors
		},
	)
}

// initDefaultRules initialize default rules
// and overrides the default rules if RuleConfig is provided by the user
func (config *RuleConfig) initDefaultRules() {

	// default config for rules
	var defaultConfig = &RuleConfig{
		Verbose: false,
		Length: Length{
			Limits: map[string]Limit{
				"Use":     {Min: 2, Max: 20},
				"Short":   {Min: 15, Max: 200},
				"Long":    {Min: 100, Max: 10000},
				"Example": {Min: 100, Max: 10000},
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
