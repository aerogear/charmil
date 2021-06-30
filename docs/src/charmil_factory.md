## Charmil Factory
The Charmil Factory offers you one-stop access to all of the Charmil packages all at once. So you just need to initialize factory in your cobra command and start using logging, i18n, etc. 
```go
import github.com/aerogear/charmil/core/factory
```

## How to use

1. Create a cobra command for your plugin
```go
func MyCommand() (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "my",
		Short: "This is my command",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	return cmd, nil
}
```

1. Create a factory instance in your command, which requires the localizer to be passed
```go
// Init Localizer
loc, err := localize.InitLocalizer(localize.Config{Language: language.English, Path: "active.en.yaml"})
if err != nil {
    return nil, err
}
// Create new/default instance of factory
newFactory := factory.Default(loc)
```

3. Now you are ready to use the packages provided by factory.

