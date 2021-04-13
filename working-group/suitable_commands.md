# Suitable commands

### help

### install - Downloads, Extract & Install a plugin

- [x] Install one NAME or multiple plugins [Name, Name ...]
- [ ] Install plugins from a newline-delimited txt file
- [ ] Install plugin from a local/URL custom manifest file
- [x] If a plugin is already installed, it will be skipped
- [x] Failiure to install a plugin will not stop the installation of other plugins
- [ ] Cleanup Stale Installations removes the versions that aren't the current version

### list

- [x] show list of all plugins
- [x] show list of installed plugins

### search

- [ ] fuzzy search in plugins via NAME, DESCRIPTION

### uninstall

- [x] uninstall a plugin NAME or list of plugins [NAME, NAME ...]
- [x] Failure to uninstall a plugin will result in an error and exit immediately

### version

- Show charmil version info & diagnostics about charmil
- GitTagrelease name
- GitCommit git revision ID
- DefaultIndexURI
- BasePath - root directory for charmil installation
- IndexPath - path where stores the local copy of the index git repository
- InstallPath - directory for plugin installations
- BinPath - directory for the symbolic links to the installed plugin executables
