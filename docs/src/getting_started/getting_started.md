---
title: Getting Started
slug: /getting_started
---

# Getting Started

This guide assumes that you already have a working Go environment, if not then please refer to
[this page](https://golang.org/doc/install) first.

## Charmil CLI

Charmil offers its own command-line interface (CLI) that allows developers to construct their new Charmil project with additional support for other templates, allowing them to focus on more essential aspects of their project.

- ### Installation:

  The Charmil CLI can be installed by any of the following ways:

  - By downloading the binary from the [releases page](https://github.com/aerogear/charmil/releases).

  - By using the `go get` command:

    ```bash
    $ go get github.com/aerogear/charmil/cli/cmd/charmil
    ```

    This will create the Charmil CLI executable under your `$GOPATH/bin` directory.

- ### Usage:
  For usage instructions on Charmil CLI, refer to [this link](../charmil_cli.md).

## Charmil Starter

The Charmil Starter Template allows users to quickly create plugins and host commands that make use of the features offered by Charmil. This template includes all of the tools and solutions needed to create a command-line tool.

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
