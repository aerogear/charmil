# Charmil Preview

The Charmil Preview package allows developers to easily add a developer preview program to their command-line interface (CLI) applications made using [Cobra](https://github.com/spf13/cobra), via feature flags.

Through this program, the developers can enable the end-users to get early access to some unreleased commands of their CLI and interact with the development team to help shape and improve the experience of other users.

## Steps to Use:

Enabling the developer preview feature through Charmil is just a simple 4-step process.

1. Open the file where the root command of your CLI is defined using Cobra.

2. Import the Charmil Preview package by adding the following line at the top of that file:

   ```go
   import "github.com/aerogear/charmil/pkg/preview"
   ```

3. Call the `InitFlag` function while passing the root command of your CLI as an argument to add a feature flag named `dev_preview` to your CLI.

   _Example:_

   ```go
   preview.InitFlag(rootCmd)
   ```

   where `rootCmd` represents the variable that holds the initialized cobra struct of your root command.

4. Call the `AddCommands` function while passing the commands of your CLI that you want to preview, as arguments.

   _Example:_

   ```go
   preview.AddCommands(subCmd1, subCmd2, subCmd3)
   ```

   where `subCmd1`, `subCmd2` and `subCmd3` represent the variables that hold the initialized cobra structs of the sub-commands that you need to preview.

&nbsp;
&nbsp;
And with that done, you're all set to use the developer preview functionality.

Now, by enabling the `dev_preview` flag, the end-users will be able to get early access to the commands of your CLI that you've added as a preview.

#### Here's an example for the same:

1. **Before:**

   ```bash
   $ ./host -h

   Host CLI for embedding commands

   Usage:
     Host [flags]
     Host [command]

   Available Commands:
     date        tell date
     help        Help about any command

   Flags:
     -d, --dev_preview   Enable dev preview commands
     -h, --help          help for Host
   ```

2. **Enabling the early access feature flag:**

   ```bash
   $ ./host --dev_preview=true
   ```

3. **After:**

   ```bash
   $ ./host -h

   Host CLI for embedding commands

   Usage:
     Host [flags]
     Host [command]

   Available Commands:
     date        tell date
     echo        [Preview] Prints given strings to stdout
     help        Help about any command

   Flags:
     -d, --dev_preview   Enable dev preview commands
     -h, --help          help for Host
   ```

   Here, `echo` is the command which was added for preview.
