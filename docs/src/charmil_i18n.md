---
title: Charmil Localization
slug: /charmil_localization
---

Localization has never been easier. Charmil provides out of the box support for the localization of your CLI.

## How to use

1. Initialize a variable using go:embed, for storing the locales directory
```go
import (
    "github.com/aerogear/charmil/cli/internal/factory"
    "github.com/aerogear/charmil/core/config"
    "github.com/aerogear/charmil/core/localize"
)

//go:embed locales/*
defaultLocales embed.FS
```
2. Provide the default language, the path to your locales, and the file format of your locales to initialize the localizer. Charmil accepts locales in the yaml, toml, and json formats.
```go

// create a config handler
h := config.NewHandler("./config.json", cfg)

// Loads config values from the local config file
err := h.Load()
if err != nil {
    log.Fatal(err)
}

// Initialize localizer providing the language, locales and format of locales file
localizer, err := localize.New(
    localize.Config{
        Language: &language.English,
        Files:    defaultLocales,
        Format:   "yaml",
    }
)
if err != nil {
    return nil, err
}

// Creates a new factory instance with default settings
cmdFactory := factory.Default(localizer, h)
```

3. LocalizeByID is a Factory function that takes a message Id stored in locales and also allows you to give templateEntries for your locales.
```go
// using Localizer for localization of cobra command
cmd := &cobra.Command{
      Use:          cmdFactory.Localizer.LocalizeByID("sample.cmd.use"),
      Short:        cmdFactory.Localizer.LocalizeByID("sample.cmd.short"),
      Example:      cmdFactory.Localizer.LocalizeByID("sample.cmd.example"),
}
```
```go
// providing template entries
cmdFactory.Localizer.LocalizeByID("sample.hi", {"Name": "John"})
```

## Sample Locales file
You must provide locales for the CLI to work. Here's a yaml-formatted sample.

```yaml
# localization by id
sample.cmd.use:
  description: "Use of sample"
  one: "sample"
sample.cmd.short:
  description: "short description of sample command"
  one: "tell sample"
sample.cmd.example:
  description: "Examples of sample command"
  one: "$ host sample"

# using templates
sample.hi:
  description: "Say hi"
  one: "Hi {{.Name}}"
```
