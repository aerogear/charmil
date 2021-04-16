# Krew-index Research

## Structure of Index

Krew's plugin index is hosted on a different repo, ie. outside its main project repo.
The repo consists of some documentation and a folder named `plugins` that stores the metadata of each plugin in separate YAML files.

## Format of manifest files

```yaml
apiVersion:
kind:
metadata:
  name:
spec:
  version:
  homepage:
  shortDescription:
  description:
  platforms:  # Specifies installation methods for various platforms (os/arch)
 - selector:
      matchExpressions:
      - key:
        operator:
        values:
    uri:  # Specifies .zip or .tar.gz archive URL of a plugin
    sha256:  # sha256sum of the url above
    files:  # Lists which files should be extracted out from downloaded archive
    - from:
      to:
    - from:
      to:
    bin:  # Specifies the path to the plugin executable among extracted files
```

## How the index is being read and processed

To understand this, let's take an example of the command:
`$ kubectl krew install whoami starboard podevents`

The following steps take place when this command is run:

- The `init()` function of the file [install.go](https://github.com/kubernetes-sigs/krew/blob/master/cmd/krew/cmd/install.go) (initialized using cobra) is called.
- All the arguments after `install` ie. `whoami`, `starboard` and `podevents` are stored in an array named `pluginNames`.
- Since the `krew install` command accepts different kinds of arguments for performing different kinds of operations as mentioned over [here](https://github.com/kubernetes-sigs/krew/blob/61868455ec219dc67b17938e16bc2e70ba1406d5/cmd/krew/cmd/install.go#L53-L66), therefore the `pluginNames` array is traversed and some checks are performed on each element to identify the type of arguments.
- After traversal, it is observed that of [all the possible types](https://github.com/kubernetes-sigs/krew/blob/61868455ec219dc67b17938e16bc2e70ba1406d5/cmd/krew/cmd/install.go#L53-L66), the command mentioned above is of the type: `kubectl krew install NAME [NAME...]`, ie. the one which directly accepts plugin names present in the **default** index as arguments.
- A slice is declared where each element will be of the type [Plugin](https://github.com/kubernetes-sigs/krew/blob/61868455ec219dc67b17938e16bc2e70ba1406d5/pkg/index/types.go#L21-L27) and will store the metadata of every plugin in the `pluginNames` array after its corresponding manifest YAML file has been parsed.
- Now the local copy of the default index is searched for manifest files of every plugin present in `pluginNames`, which are parsed into variables of type [Plugin](https://github.com/kubernetes-sigs/krew/blob/61868455ec219dc67b17938e16bc2e70ba1406d5/pkg/index/types.go#L21-L27) and then appended to the slice mentioned in the last point.
- Finally, every element of the slice mentioned above is passed as an argument into the [Install](https://github.com/kubernetes-sigs/krew/blob/61868455ec219dc67b17938e16bc2e70ba1406d5/internal/installation/install.go#L58) function. This function takes metadata of plugins as arguments and installs them into the system.

## Process to add new plugin

We just need to tag a git release of the plugin archive with a [semantic version](https://semver.org/) (e.g. `v1.0.0`) and then create a PR to the [krew-index](https://sigs.k8s.io/krew-index) repository with our plugin manifest file (e.g. `my-plugin.yaml`) to the `plugins/` directory.

Once this is done, the plugin will be available for installation through Krew.

## Process to update a plugin

There are 2 ways in which a plugin can be updated:

- **Manually:**
  - Update the `version`, `uri`, and `sha256` fields of the plugin manifest file and send a PR to the [krew-index](https://sigs.k8s.io/krew-index) repo.
  - If the PR introduces updates to only the above 3 fields, then it will be merged automatically with the help of [this](https://github.com/ahmetb/krew-index-autoapprove) bot. See an example [here](https://github.com/kubernetes-sigs/krew-index/pull/508).
- **Automated:**
  - The [krew-release-bot](https://github.com/rajatjindal/krew-release-bot) is a Github Action which automatically bumps the version in `krew-index` repo every time we push a new git tag to our repository and creates a PR on our behalf.
  - [Here](https://github.com/kubernetes-sigs/krew-index/pull/490) is an example of a completely automated update process using both the bots mentioned above.
