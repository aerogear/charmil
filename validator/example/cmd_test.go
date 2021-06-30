package example

import (
	"testing"

	"github.com/aerogear/charmil/validator/rules"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewCommand()

	// Testing cobra commands with default recommended config
	// default config can also be overrided
	var vali rules.RuleConfig = rules.RuleConfig{
		// Verbose: true,
	}
	validationErr := vali.ExecuteRules(cmd)
	for _, errs := range validationErr {
		if errs.Err != nil {
			t.Errorf("%s: cmd %s: %s", errs.Rule, errs.Cmd.CommandPath(), errs.Name)
		}
	}

	if len(validationErr) == 0 {
		t.Log("success, Test Passed!")
	}

}
