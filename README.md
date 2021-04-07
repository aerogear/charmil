# cobra-cmd-hub

## Deliverables

- Use a CLI plugin as an extension in some CLI
- Link multiple CLI's together
- Registry Manager containing the list and metadata of CLI's
- Github repo/pages based remote index servers for plugins which will contain docs, metadata of the plugin
- Framework that will be able to read these remote index servers

## Registery CLI

We can build a CLI that will be installed in user's system which would perform the below tasks.

![Registry CLI](mockups/cli-registry.png)

## Specification for the CLI registry

```json
{
  "name": "",
  "description": "",
  "version": "",
  "environment": "",
  "registry": "",
  "homepage": "",
  "bin": {},
  "private": false,
  "dependencies": {},
  "maintainers": {},
  "license": ""
}
```

## Registry Manager

This is the place where developers will be able to push their CLI's as extensions. All the information about the extensions like name, versions, docs, homepage, licence, size, repo, etc will be present.

We can use GitHub packages to host packages and use them as extensions/dependencies in other CLI's. It will also help developers to create DevOps end to end workflow.

## Executing Extensions in host CLI using Registry CLI

This can have two modes -

1. Local Mode - Package can be installed that are already present in the $GOPATH
2. Global Mode - Package can be installed from Registry manager.

Package (go CLI module) will be added to `specification.json` & `go.mod` file of project and hence can be used as features in the host CLI.

Feature to remove and update package in Registry CLI
