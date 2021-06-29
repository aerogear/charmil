package rules

import (
	"encoding/json"
	"fmt"
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

	lenErrs := validateLength(cmd, config.Length, config.Verbose)
	info.TotalErrors += len(lenErrs)

	mustExistErrs := validateMustExist(cmd, config.MustExist, config.Verbose)
	info.TotalErrors += len(mustExistErrs)

	info.Errors = append(info.Errors, lenErrs...)
	info.Errors = append(info.Errors, mustExistErrs...)
	info.TotalTested++

	for _, child := range cmd.Commands() {

		// base case
		if !child.IsAvailableCommand() || child.IsAdditionalHelpTopicCommand() {
			continue
		}

		// recursive call
		if err := config.executeHelper(child, info); err != nil {
			return err
		}
	}

	// prints additional info in debug mode
	if config.Verbose {
		fmt.Printf("commands checked: %d\nchecks failed: %d\n", info.TotalTested, info.TotalErrors)
	}

	return info.Errors
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
