package main

import (
	"testing"

	"github.com/aerogear/charmil/validator/rules"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewCommand()

	// Testing cobra commands with default recommended config
	// default config can also be overrided
	ruleCfg := rules.ValidatorConfig{
		ValidatorOptions: rules.ValidatorOptions{
			IgnoreCommands: map[string]bool{"cmd0": true},
		},
		ValidatorRules: rules.ValidatorRules{
			Length: rules.Length{
				Limits: map[string]rules.Limit{
					"Use": {Min: 1},
				},
			},
			MustExist:  rules.MustExist{Fields: map[string]bool{"Run": true}},
			UseMatches: rules.UseMatches{Regexp: `^[^-_+]+$`},
		},
	}

	validationErr := rules.ExecuteRules(cmd, &ruleCfg)
	if len(validationErr) != 0 {
		t.Errorf("validationErr was not empty, got length: %d; want %d", len(validationErr), 0)
	}
	for _, errs := range validationErr {
		if errs.Err != nil {
			t.Errorf("%s: cmd %s: %s", errs.Rule, errs.Cmd.CommandPath(), errs.Name)
		}
	}

}
