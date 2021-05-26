## Key differences between Krew and Charmil

|            -             |                   **Krew**                    | **Charmil**                                                            |
| :----------------------: | :-------------------------------------------: | ---------------------------------------------------------------------- |
| _Role of the developer:_ |        To develop a CLI and publish it        | To link different CLIs to their own CLI and then publish the final CLI |
| _Role of the end-user:_  | To link different available CLIs and use them | To directly use any linked CLI published by developers                 |

## Personas that Charmil will cover

- ### SDK users (aka CLI developers):

  - As mentioned in the table above, their role would be to use our SDK to bring some extensions into their own CLI and make it available for the end-users of Charmil to use.
  - They have an option to add multiple custom indexes.

- ### CLI end-users:
  - For them, Charmil will just be like a simple package manager (similar to [npm](https://www.npmjs.com/)), through which they can manage (install, upgrade, etc.) the CLIs developed by Charmil SDK users.
  - They will have access to commands such as: `install`, `list`, `search`, `upgrade`, `uninstall` and `version`.

## An example to sum up the entire model

Let's assume that with the help of our SDK, a developer wants to create a CLI named `A`, that uses the CLIs named `B`, `C` and `D` in the background.

To do so, the following steps can be followed:

1. The developer can start by adding a custom index that contains the metadata of CLIs: `B`, `C` and `D`.
2. Now our SDK will use that custom index and assist the developer in creating a new CLI: `A` using the other 3 CLIs ( ie. `B`, `C`, `D`).
3. Once the CLI: `A` has been created successfully, the developer can publish it using 'GitHub Packages' and Charmil will make it available for people to use.
4. If the developer wants to release new updates for users in the future, then it can simply be done by pushing the new release of the CLI package on GitHub and tagging it with a [semantic version](https://semver.org/) (e.g. `v1.1.0`) and Charmil will handle the rest.

Now people can use the Charmil CLI to install CLI: `A`, by running the command:

```python
charmil install A
```

Similarly Charmil will allow end-users to upgrade and uninstall the package by using the following commands:

```python
charmil upgrade A
```

```python
charmil uninstall A
```
