---
title: Charmil Localization
slug: /charmil_localization
---

Localization has never been easier. Charmil provides out of the box support for the localization of your CLI.

## How to use

1. Provide the default language, the path to your locales, and the file format of your locales to initialize the localizer. Charmil accepts locales in the yaml, toml, and json formats.
```go
// Initialize localizer providing the language, locales and format of locales file
loc, err := localize.InitLocalizer(localize.Config{Language: language.English, Path: "examples/plugins/date/locales/en/en.yaml", Format: "yaml"})
if err != nil {
    return nil, err
}

// Create new/default instance of factory
newFactory := factory.Default(loc)
```

2. LocalizeByID is a Factory function that takes a message Id stored in locales and also allows you to give templateEntries for your locales.
```go
// using Localizer for localization of cobra command
cmd := &cobra.Command{
    Use:          newFactory.Localize.LocalizeByID("sample.cmd.use"),
    Short:        newFactory.Localize.LocalizeByID("sample.cmd.short"),
    Example:      newFactory.Localize.LocalizeByID("sample.cmd.example"),
}
```
```go
// providing template entries
newFactory.Localize.LocalizeByID("sample.hi", {"Name": "John"})
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
