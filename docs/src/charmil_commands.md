## Charmil Commands
Charmil provides prebuilt commands for developers to use when testing their application. These commands are opt-in and are only added when the application is built with a specific tag.

To add charmil commands to your root command you can call the `CharmilCommands()` function. If calling this function does not populate `err` then add the commands to your root command.  

```go
cmd := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
	}
	charmilCommands, err := commands.CharmilCommands()
	if err == nil {
		cmd.AddCommand(charmilCommands)
	}
```
The reason for this structure is because sometimes you do not want to add charmil commands to your root command, for example in a production build. So that is why `CharmilCommands()` populating err does not mean it is time to panic.

Charmil commands are opt-in meaning you must specifically ask for them in your build and if you don't trying to get 
them will return an error. To not get an error you must build your application; in this example `ack` with the `dev` tag.
```bash
go build -tags dev ack.go
```
Now this build will make charmil commands available to use. You can run a charmil command by running the `charmil` command and any subcommand.
```bash
ack charmil validate
```