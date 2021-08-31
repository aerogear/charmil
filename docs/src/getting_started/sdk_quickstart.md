---
title: Charmil SDK Quickstart
slug: /sdk_quickstart
---

# Charmil SDK Quickstart

This is the component responsible for constructing modular, multi-repo Golang CLIs with Cobra while providing a variety of high-level solutions for typical challenges like configuration, internationalization, etc.

- ## Installation:

  First use `go get` to install the latest tagged release of the framework.
  The following command will install the framework along with its dependencies:

  ```bash
  $ go get -u github.com/aerogear/charmil
  ```

  Next, import the required package into your project:

  - For the config management package:

    ```go
    import "github.com/aerogear/charmil/core/config"
    ```

  - For utils packages (eg. color, localize, logging, etc.):

    ```go
    import "github.com/aerogear/charmil/core/utils/{name_of_the_required_package}"
    ```

- ## Usage:

  Links to usage docs:

  - [Config](../charmil_config.md)
  - [Utils](../utils)
