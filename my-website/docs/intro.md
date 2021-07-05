---
title: Charmil Commands
---
Charmil provides prebuilt commands for developers to use when testing their application. These commands are opt-in and are only added when the application is built with a specific tag.

To add charmil commands to your root command you can call the `AttachCharmilCommands()` function with your root command. It also must be noted that just not being in a dev build shouldn't cause `err` to be populated.  

```go
cmd := &cobra.Command{
		Use:          "Host",
		Short:        "Host CLI for embedding commands",
		SilenceUsage: true,
}

err = commands.AttachCharmilCommands(cmd)
if err != nil {
	log.Fatal(err)
}
```

Charmil commands are opt-in meaning you must specifically ask for them in your build; meaning trying to call `AttachCharmilCommands()` does not do anything in these builds.
```bash
go build -tags dev ack.go
```
Now this build will make charmil commands available to use. You can run a charmil command by running the `charmil` command and any subcommand.
```bash
ack charmil validate
```