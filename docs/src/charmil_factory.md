---
title: Charmil Factory
slug: /charmil_factory
---

The Charmil Factory offers you one-stop access to all of the Charmil packages all at once. So you just need to initialize factory in your cobra command and start using logging, i18n, etc.

```go
import (
    "<module-name>/<cli-name>/pkg/factory"
    "github.com/aerogear/charmil/core/config"
    "github.com/aerogear/charmil/core/utils/localize"
)
```

## How to use

1. Initialize factory in your command, by providing it an instance of charmil localizer and config handler.

```go
// Embed provides access to files embedded in the running Go program meaning you can get your locales directory with the path to it.
//go:embed locales/*
defaultLocales embed.FS

// create a config handler
h := config.NewHandler("./config.json", cfg)

// Initialize localizer providing the language, locales and format of locales file
localizer, err := localize.New(
    localize.Config{
        Language: &language.English,
        Files:    defaultLocales,
        Format:   "yaml", // charmil accepts locales in the yaml, toml, and json formats.
    }
)
if err != nil {
    return nil, err
}

// Creates a new factory instance with default settings
cmdFactory := factory.Default(localizer, h)
```

2. Now you are ready to use the packages/utilities provided by factory.
   - [Charmil i18n](./utils/charmil_i18n.md)
   - [Charmil Config](./charmil_config.md)
   - [Charmil Logger](./utils/charmil_logger.md)
