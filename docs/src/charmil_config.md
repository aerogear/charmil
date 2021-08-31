---
title: Charmil Config Management Package
slug: /charmil_config
---

# Charmil Config Management Package

The Charmil Config package offers a convenient mechanism for both host and plugin developers to manage configurations in their command-line interface (CLI) applications made using [Cobra](https://github.com/spf13/cobra).

## Features:

- Helps in maintaining all available configurations in a single, centralized local config file.
- Provides the plugin developers with a functionality to add their CLI configurations to the host CLI local config file with ease.
- Provides the host CLI developers with a set of methods to read/write configurations from/to a local config file.

## Steps to Use:

- ### For host CLI developers:

  1.  Open the file where the root command of your CLI is defined using Cobra.

  2.  Import the Charmil Config package by adding the following line at the top of that file:
      ```go
      import c "github.com/aerogear/charmil/core/config"
      ```
  3.  Define a struct whose fields represent the keys to all the values that you want to store as config and create an instance of it.

      **Important**: Every field of the defined struct needs to be exportable (ie. start with an uppercase letter). The unexportable fields cannot be stored as config.

      _Example:_

      ```go
      type config struct {
        Key1 string
        Key2 string
        Key3 string
        Key4 string
      }

      cfg = &config{}
      ```

  4.  Store a new instance of the Charmil Config handler by calling the `NewHandler` function while passing the path of local config file and the instance of the config struct (initialized in the last step) as arguments.

      _Example:_

      ```go
      h = c.NewHandler("./examples/host/config.json", cfg)
      ```

  5.  Load config values from the local config file using the `Load` method.

      _Example:_

      ```go
      err := h.Load()
      if err != nil {
      	log.Fatal(err)
      }
      ```

  6.  You can set/get/modify values under any key of config using the idiomatic way to interact with structs in Golang.

      _Example:_

      ```go
      // Sets a value into config
      cfg.Key4 = "val4"

      // Overwrites a value in config
      cfg.Key2 = "newVal2"

      // Returns the value under specified key in config
      fmt.Println(cfg.Key3) // Prints: val3
      ```

  7.  Write current config into the local config file using the `Save` method.

      _Example:_

      ```go
      err = h.Save()
      if err != nil {
        log.Fatal(err)
      }
      ```

- ### For plugin developers:

  1.  Open the file where the root command of your CLI is defined using Cobra.

  2.  Import the Charmil Config package by adding the following line at the top of that file:
      ```go
      import c "github.com/aerogear/charmil/core/config"
      ```
  3.  Define a struct whose fields represent the keys to all the values that you want to store as config and create an instance of it.
      **Important**: Every field of the defined struct needs to be exportable (ie. start with an uppercase letter). The unexportable fields cannot be stored as config.

      _Example:_

      ```go
      type config struct {
        Key5 string
        Key6 string
        Key7 string
        Key8 string
      }

      cfg = &config{}
      ```

  4.  You can set/get/modify values under any key of config using the idiomatic way to interact with structs in Golang.

      _Example:_

      ```go
      // Sets values into config
      cfg.Key5 = "val5"
      cfg.Key6 = "val6"
      cfg.Key7 = "oldVal7"
      cfg.Key8 = "val8"

      // Overwrites a value in config
      cfg.Key7 = "val7"

      // Returns the value under specified key in config
      fmt.Println(cfg.Key6) // Prints: val6
      ```

  5.  Use the `MergePluginCfg` function to merge the current plugin config into the host CLI config struct.

      _Example:_

      ```go
      err = c.MergePluginCfg(pluginName, h, cfg)
      if err != nil {
      	log.Fatal(err)
      }
      ```

      where `pluginName` is the name of the plugin (as you want it in the local config file), `h` is the config handler passed from the host CLI and `cfg` is a pointer to an instance of the current config file (initialized in step 3).

## Here's an example for the same:

- #### Initial Configurations [Before running the Host CLI]:

  - _`config.json`_ file:

    ```json
    {
      "key1": "val1",
      "key2": "val2",
      "key3": "val3"
    }
    ```

  - Plugin A's Config Struct:
    ```go
    {val4 val5 val6}
    ```
  - Plugin B's Config Struct:
    ```go
    {val7 val8 val9}
    ```

- #### Final Configurations [After running the Host CLI]:
  - _`config.json`_ file:
    ```json
    {
      "key1": "val1",
      "key2": "val2",
      "key3": "val3",
      "plugins": {
        "pluginA": {
          "Key4": "val4",
          "Key5": "val5",
          "Key6": "val6"
        },
        "pluginB": {
          "Key7": "val7",
          "Key8": "val8",
          "Key9": "val9"
        }
      }
    }
    ```
