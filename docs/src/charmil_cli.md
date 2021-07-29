---
title: Charmil CLI
slug: /charmil_cli
---

> Note: Charmil CLI is under development

Charmil CLI provides a way to create and manage the Charmil project. CLI gives you preconfigured shell commands as recommended by Charmil, so you don't have to integrate charmil core yourself. Available commands are:

## Init Command

Init command initializes a Charmil project boilerplate(starter). This will generate shell commands, readme, licence, etc to get you started easily.

```bash
$ charmil init
```

Fill in the prompted info(github owner, repository and cli name) to get started!

## Add Command

Add command adds a new command into the CLI

```bash
$ charmil add CMD_NAME
```

## Crud Command

With the help of `crud` command, developers can eliminate a lot of boilerplate in CLIs containing multiple services that perform standard CRUD operations.

Using a set of pre-defined templates, this command generates CRUD files in the directory specified with the `path` flag.

These generated files can be modified by developers to fit their own needs.

### Usage:

```bash
charmil crud [flags]
```

### Example:

```bash
$ charmil crud --singular=kafka --plural=kafkas --crudpath="./kafka" --localepath="./cmd/locales/en"
```

### Flags:

```
  -c, --crudpath string     path where CRUD files need to be generated (default ".")
  -h, --help                help for crud
  -l, --localepath string   path where the language file needs to be generated (default ".")
  -p, --plural string       name in plural form (REQUIRED)
  -s, --singular string     name in singular form (REQUIRED)
```
