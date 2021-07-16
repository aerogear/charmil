package echo

import (
	"testing"

	c "github.com/aerogear/charmil/core/config"
	"github.com/aerogear/charmil/core/factory"
	"github.com/aerogear/charmil/validator"
	"github.com/aerogear/charmil/validator/rules"
)

func Test_ExecuteCommand(t *testing.T) {
	cfg := &config{}
	h := c.NewHandler("./examples/plugins/echo/test_config.json", cfg)
	cmdFactory := factory.Default(nil, h)

	cmd, err := EchoCommand(cmdFactory)
	if err != nil {
		t.Fatal(err)
	}

	// Testing cobra commands with default recommended config
	// default config can also be overrided
	ruleCfg := rules.ValidatorConfig{
		ValidatorOptions: rules.ValidatorOptions{
			SkipCommands: map[string]bool{"cmd100 cmd0 subcmd01": true},
		},
		ValidatorRules: rules.ValidatorRules{
			Length: rules.Length{
				RuleOptions: validator.RuleOptions{
					SkipCommands: map[string]bool{"cmd100": true},
				},
				Limits: map[string]rules.Limit{
					"Use": {Min: 1},
				},
			},
			MustExist: rules.MustExist{
				RuleOptions: validator.RuleOptions{
					SkipCommands: map[string]bool{"cmd100": true},
				},
				Fields: map[string]bool{"Run": true},
			},
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
