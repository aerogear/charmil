# Wrapper POC
Provide wrapper structure without any strong dependency on Cobra. Plugin Core based of those structures build cobra command.

## Architecture & Working
1. An Host CLI in which commands need to be embedded
2. A Minimal YAML file defining the commands of the plugin CLI.

    ```yaml
    commands:
    - 
        name: COMMAND_NAME
        mapsTo:
            name: INITIAL_NAME
            subcommand: SUBCOMMAND_NAME
            args:
                - ARGUMENT_1
                - ARGUMENT_2
        shortDescription: "description of command"
        usage: '
        Explain the usage of the command in host CLI
        $ HOST_NAME COMMAND_NAME ARGUMENT_1
        $ HOST_NAME COMMAND_NAME ARGUMENT_2
        '
        flags:
            # define flags
    ```
3. Plugin Structure to extract the plugin commands, args, flags etc. defined in YAML file
4. Plugin Loader to load the YAML and create cobra command in the Host (provides common API for host and extensions to use)

## pros
- Access to huge market. We can onboard all the CLI's into this framework no matter they are made with golang, bash, python etc.
- Easy onboarding of plugins. A minimal configuration yaml file to be made, providing the details about using the commands, arguments and flags.

## cons
- Wrapping the binaries may give performance issues
- Assuming that user's system has CLI installed which is to be embedded (we should be able to solve this)
