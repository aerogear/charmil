---
title: Charmil Example Builder
slug: /charmil_example_builder
---

## Example Builder

Example Builder provides Example string builder which generates full example string of the cobra command and append to existing examples.

NewCmdExample function accepts a cobra command in which example needs to be generated, description of the command and a slice of flags(if any)

```go
examplebuilder.NewCmdExample(cmd, "List all artifacts for the default artifacts group", []string{"list --page=4"})
examplebuilder.NewCmdExample(cmd, "List all artifacts", []string{"list"})
```

Output

```bash
Examples:

# List all artifacts for the default artifacts group
kafka list --page=4

# List all artifacts
kafka list
```
