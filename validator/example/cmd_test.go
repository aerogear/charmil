package main

import (
	"testing"

	"github.com/aerogear/charmil/validator/rules"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewCommand()

	// Testing cobra commands with default recommended config
	// default config can also be overrided
	var vali rules.RuleConfig = rules.RuleConfig{
		Rules: []rules.Rules{
			&rules.Length{Verbose: true, Limits: map[string]rules.Limit{"Use": {Min: 100}}},
		},
	}

	validationErr := vali.ExecuteRules(cmd)
	if len(validationErr) != 0 {
		t.Errorf("validationErr was not empty, got length: %d; want %d", len(validationErr), 0)
	}

	for _, errs := range validationErr {
		if errs.Err != nil {
			t.Errorf("%s: cmd %s: %s", errs.Rule, errs.Cmd.CommandPath(), errs.Name)
		}
	}

}
