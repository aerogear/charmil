## Must Exist Rule
This rule ensures that the fields specified in the rule must be present in the cobra command.

## How to use
1. Create a MustExist rule by providing the fields you want to be present in cobra command
```go
myExistRule := &rules.MustPresent{
    Fields: []string{"Use", "Short", "Long", "Example", "SilenceUsage", "PreRun", "Hi"},
}
``` 
2. Use ValidateMustPresent function to validate your cobra command with the rule created above. It takes a verbose bool for enabling/disabling debug StatusLogs
```go
errr := myExistRule.ValidateMustPresent(cmd, true)
```
3. ValidateMustPresent function will return a slice of `ValidationError` type, which will allow you to efficiently test your code