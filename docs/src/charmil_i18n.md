---
title: Charmil Localization
slug: /charmil_localization
---

Localization has never been easier. Charmil provides out of the box support for the localization of your CLI.

## How to use

1. Initialize factory in your command. Follow the [Factory documentation](./charmil_factory.md)
2. LocalizeByID is a Factory function that takes a message Id stored in locales and also allows you to give templateEntries for your locales.
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
