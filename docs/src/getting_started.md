---
title: Getting Started Guide
slug: /getting_started
---

# Getting Started with Charmil

This guide assumes that you already have a working Go environment, if not then please refer to
[this page](https://golang.org/doc/install) first.

## Charmil CLI

Charmil offers its own command-line interface (CLI) that allows developers to construct their new Charmil project with additional support for other templates, allowing them to focus on more essential aspects of their project.
&nbsp;

- ### Installation:

  In order to use the Charmil CLI, run the following command:

  ```bash
  $ go get github.com/aerogear/charmil/cli/cmd/charmil
  ```

  This will create the Charmil CLI executable under your `$GOPATH/bin` directory.

- ### Usage:
  For usage instructions on Charmil CLI, refer to [this link](./charmil_cli.md).

## Charmil Starter

The Charmil Starter Template allows users to quickly create plugins and host commands that make use of the features offered by Charmil. This template includes all of the tools and solutions needed to create a command-line utility.
&nbsp;

- ### Usage:

  The Charmil Starter Template can be used by any of the following ways:

  - By using the `charmil init` command offered by the Charmil CLI
    ```bash
    $ charmil init
    ```
  - By cloning the [Charmil Starter Github repository](https://github.com/aerogear/charmil-starter)
    ```bash
    $ git clone https://github.com/aerogear/charmil-starter.git
    ```

## Charmil SDK

This is the component responsible for constructing modular, multi-repo Golang CLIs with Cobra while providing a variety of high-level solutions for typical challenges like configuration, internationalization, etc.
&nbsp;

- ### Installation:

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

- ### Usage:

  Links to usage docs:

  - [Config](./charmil_config.md)
  - [Utils](./utils)

## Charmil Validator

- ### Installation:

  First use `go get` to install the latest tagged release of the framework.
  The following command will install the framework along with its dependencies:

  ```bash
  $ go get -u github.com/aerogear/charmil
  ```

  Next, import the Charmil Validator package into your project:

  ```go
  import "github.com/aerogear/charmil/validator"
  ```

- ### Usage:

  For usage instructions on Charmil Validator, please refer to [this link](./charmil_validator.md#how-to-use).
