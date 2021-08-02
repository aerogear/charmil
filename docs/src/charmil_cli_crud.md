# Charmil CRUD Generator

- With the help of the `charmil crud` command, developers can eliminate a lot of boilerplate in CLIs containing multiple services that perform standard CRUD operations.

- Using a set of pre-defined templates, this command generates CRUD command packages in the directory specified with the `crudpath` flag and the corresponding language file in the directory specified with the `localepath` flag.

- These generated files can be modified by developers to fit their own needs.

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

## Example:

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
     ┗ 📜root.go
  ```
