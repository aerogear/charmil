package main

import (
	"testing"

	"github.com/aerogear/charmil/validator/rules"
)

func Test_ValidateCommandsUsingCharmilValidator(t *testing.T) {
	cmd := Root()

	// Testing cobra commands with default recommended config
	vali := rules.ValidatorConfig{
		ValidatorRules: rules.ValidatorRules{
			Length: rules.Length{
				Limits: map[string]rules.Limit{
					"Short":   {Min: 5},
					"Example": {Min: 10},
					"Long":    {Min: 10},
				},
			},
		},
	}
	validationErr := rules.ExecuteRules(cmd, &vali)

	if len(validationErr) != 0 {
		t.Errorf("validationErr was not empty, got length: %d; want %d", len(validationErr), 0)
	}

	for _, errs := range validationErr {
		if errs.Err != nil {
			t.Logf("%s: cmd %s: %s", errs.Rule, errs.Cmd.CommandPath(), errs.Name)
		}
	}
}
