---
title: Charmil CRUD Generator Command
slug: /charmil_cli_crud
---

# Charmil CRUD Generator

- With the help of the `charmil crud` command, developers can eliminate a lot of boilerplate in CLIs containing multiple services that perform standard CRUD operations.

- Using a set of pre-defined templates, this command generates CRUD command packages in the directory specified by the `crudpath` flag as well as its corresponding language file in the directory specified by the `localepath` flag.

- These generated files can then be modified by developers to fit their own needs.

## Usage:

```bash
charmil crud [flags]
```

## Flags:

```
  -c, --crudpath string     path where CRUD files need to be generated (default ".")
  -h, --help                help for crud
  -l, --localepath string   path where the language file needs to be generated (default ".")
  -p, --plural string       name in plural form (REQUIRED)
  -s, --singular string     name in singular form (REQUIRED)
```

## Steps to use:

- Let's say you need to generate CRUD commands for managing your Kafka instances, the following command can be used for the same:

  ```bash
  $ charmil crud --singular=kafka --plural=kafkas --crudpath="./kafka" --localepath="./cmd/locales/en"
  ```

- On running the command mentioned above, the required files will be generated in your project in the following structure:

  ```code
  📦Your CLI
   ┣ 📂cmd
   ┃ ┗ 📂locales
   ┃   ┗ 📂en
   ┃     ┗ 📜crud.en.yaml
   ┗ 📂kafka
     ┣ 📂create
     ┃ ┣ 📜create.go
     ┃ ┗ 📜run.go
     ┣ 📂delete
     ┃ ┣ 📜delete.go
     ┃ ┗ 📜run.go
     ┣ 📂describe
     ┃ ┣ 📜describe.go
     ┃ ┗ 📜run.go
     ┣ 📂list
     ┃ ┣ 📜list.go
     ┃ ┗ 📜run.go
     ┣ 📂use
     ┃ ┣ 📜use.go
     ┃ ┗ 📜run.go
     ┗ 📜kafka.go
  ```

- Once the CRUD packages have been generated, go to the generated `kafka.go` file and add all the missing imports there.

- Using the following line, add the generated CRUD commands to your CLI:

  ```go
  cmd.AddCommand(kafka.NewCommand(cmdFactory))
  ```

  where `cmd` refers to your CLI's parent command and `cmdFactory` refers to the factory instance.

Now you're all set to use the CRUD commands in your CLI.
