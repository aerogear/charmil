## Length Rule
The lengths of properties in string format provided by `cobra.Command` can be controlled using the Length Rule.

## How to use
1. Create a length rule by providing the limits for each attribute you want to control
```go
myLengthRule := &rules.Length{
    Limits: map[string]rules.Limit{
        "Use":   {Min: 1, Max: 5},
        "Short": {Min: 4, Max: 10},
        "Long":  {Min: 5, Max: 20},
        "Not":   {Min: 2, Max: 22},
    },
}
```
2. Use ValidateLength function to validate your cobra command with the rule created above. It takes a verbose bool for enabling/disabling debug StatusLogs
```go
err := myLengthRule.ValidateLength(cmd, true)
```
3. ValidateLength function will return a slice of `ValidationError` type, which will allow you to efficiently test your code