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

Add command adds a new command into the CLI, along with english locales for the command.

```bash
$ charmil add --cmdName="mycmd"
```
if you want to specify the path for command to be created, use `--cmdPath` flag. By default `--cmdPath` is set to current directory.
```bash
$ charmil add --cmdName="mycmd" --cmdPath="./cmd"
```


## Crud Command

Helps developers generate CRUD commands for their CLI.

- #### [Link to detailed documentation on the Charmil CRUD Command](./charmil_cli_crud.md)
