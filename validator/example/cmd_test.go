package example

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/aerogear/charmil/validator/rules"
)

func Test_ExecuteCommand(t *testing.T) {
	cmd := NewCommand()

	// Testing cobra commands with default recommended config
	// default config can also be overrided by
	/*
		var vali rules.RuleConfig = rules.RuleConfig{
				Verbose: true,
				MustExist: rules.MustExist{
					Fields: []string{"Args"},
				},
			}
	*/
	var vali rules.RuleConfig
	validationErr := vali.ExecuteRules(cmd)
	for _, errs := range validationErr {
		if errs.Err != nil {
			t.Errorf("%s: cmd %s: %s", errs.Rule, errs.Cmd.CommandPath(), errs.Name)
		}
	}

	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.Execute()
	_, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
}
