# Charmil Config

The Charmil Config package offers a convenient mechanism for both host and plugin developers to manage configurations in their command-line interface (CLI) applications made using [Cobra](https://github.com/spf13/cobra).

## Features:

- Uses [Viper](https://github.com/spf13/viper) under the hood.
- Helps in maintaining all available configurations in a single central-local config file.
- Provides the plugin developers with a set of methods to maintain a configuration map and export it to the host CLI.
- Provides the host CLI developers with a set of methods to import plugin configuration maps and read/write configurations from/to a local config file.

## Steps to Use:

- ### For host CLI developers:

  1.  Open the file where the root command of your CLI is defined using Cobra.

  2.  Import the Charmil Config package by adding the following line at the top of that file:
      ```go
      import "github.com/aerogear/charmil/core/config"
      ```
  3.  Store a new instance of the Charmil Config handler by calling the `New` function.

      _Example:_

      ```go
      h := config.New()
      ```

  4.  Create an instance of the `File` struct, passing all the required settings for the local config file as fields. Once this is done, pass that `File` instance as an argument into the `InitFile` method, in order to link the local config file to the pointer receiver ie. handler.

      _Example:_

      ```go
      f = config.File{
      	Name: "config",
      	Type: "yaml",
      	Path: "./examples/host",
      }

      h.InitFile(f)
      ```

  5.  Load config values from the local config file using the `Load` method.

      _Example:_

      ```go
      err := h.Load()
      if err != nil {
      	log.Fatal(err)
      }
      ```

  6.  Use the `SetValue` and `GetValue` methods respectively to set and get values to/from the current config.

      _Example:_

      ```go
      h.SetValue("key0", "val0")

      val, err := h.GetValue("key0") // returns "val0"
      if err != nil {
      	log.Fatal(err)
      }
      ```

  7.  Map a plugin name to its config map using the `SetPluginCfg` method by passing the name of the plugin (as you want it in the config file) as the first argument and the config map imported from the plugin, as the second argument.

      _Example:_

      ```go
      h.SetPluginCfg("pluginName", pluginCfg)
      ```

  8.  Call the `MergePluginCfg` method to store the config of every plugin (ie. mapped using the `SetPluginCfg` method) into the current config.

      _Example:_

      ```go
      h.MergePluginCfg()
      ```

  9.  Write current config into the local config file using the `Save` method.

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
      import "github.com/aerogear/charmil/core/config"
      ```
  3.  Store a new instance of the Charmil Config handler by calling the `New` function.

      _Example:_

      ```go
      h := config.New()
      ```

  4.  Use the `SetValue` and `GetValue` methods respectively to set and get values to/from the current config.

      _Example:_

      ```go
      h.SetValue("key0", "val0")

      val, err := h.GetValue("key0") // returns "val0"
      if err != nil {
      	log.Fatal(err)
      }
      ```

  5.  Use the `GetAllSettings` to get the current config in the form of a map and export it to the host CLI.

      _Example:_

      ```go
      cfg := h.GetAllSettings()

      return cfg
      ```

## Here's an example for the same:

- #### Initial Configurations [Before running the Host CLI]:

  - _`config.yaml`_ file:

    ```yaml
    key1: "val1"
    key2: "val2"
    key3: "val3"
    ```

  - Plugin A's Config Map:
    ```go
    map[key4:value4 key5:value5 key6:value6]
    ```
  - Plugin B's Config Map:
    ```go
    map[key7:value7 key8:value8 key9:value9]
    ```

- #### Final Configurations [After running the Host CLI]:
  - _`config.yaml`_ file:
    ```yaml
    key1: "val1"
    key2: "val2"
    key3: "val3"
    plugins:
      pluginA:
        key4: "val4"
        key5: "val5"
        key6: "val6"
      pluginB:
        key7: "val7"
        key8: "val8"
        key9: "val9"
    ```
