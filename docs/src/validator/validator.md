## Validator
Validator can be used for testing and controlling many aspects of cobra commands. It provides many rules out of the box for validating the commands.

## Rules provided by validator
- [LengthRule](validator_length_rule.md)
- [MustExistRule](validator_must_exist_rule.md)
- UseMatches
> We are working on the validator library to provide more rules

## How to use
It is recommended to use the validator while writing unit tests for cobra commands.

1. Create a validator of type `rules.RuleConfig`. You can provide your own RulesConfig or use the default one by leaving it empty
```go
var vali rules.RuleConfig
```
or overriding default config
```go
vali := rules.RuleConfig{
	Verbose: true,
	MustExist: rules.MustExist{
		Fields: []string{"Args"},
	},
}
```
2. Generate the validation errors by using `ExecuteRules` function over the config
```go
validationErr := vali.ExecuteRules(cmd)
```
`ExecuteRules` function will return a slice of `ValidationError` object, which can be efficiently used for testing purposes.
```go
for _, errs := range validationErr {
	if errs.Err != nil {
		t.Errorf("%s: cmd %s: %s", errs.Rule, errs.Cmd.CommandPath(), errs.Name)
	}
}
```