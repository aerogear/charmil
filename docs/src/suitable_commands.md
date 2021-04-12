# Suitable commands

### help

### index

- [x] Manage(list, add, delete) custom plugin indexes

### info

- [x] Show information about a plugin via NAME & INDEX/NAME

### install - Downloads, Extract & Install a plugin

- [x] Install one NAME or multiple plugins [Name, Name] from default index
- [ ] Install plugins from a newline-delimited txt file
- [x] Install one or mutiple plugins from custom index
- [ ] Install plugin from a local/URL custom manifest file
- [x] If a plugin is already installed, it will be skipped
- [x] Failiure to install a plugin will not stop the installation of other plugins
- [ ] Cleanup Stale Installations removes the versions that aren't the current version

### list

- [x] show list of all charmil plugins
- [x] show list of installed plugins

### search

- [ ] fuzzy search in plugins via NAME, DESCRIPTION

### uninstall

- [x] uninstall a plugin or list of plugins
- [x] Failure to uninstall a plugin will result in an error and exit immediately
- uninstall command does not support INDEX/PLUGIN syntax just specify PLUGIN

### update

- [x] update local copy of plugin indexes

### upgrade

- [x] Install a plugin or all plugins to the newer version

### version

- Show charmil version info & diagnostics about charmil
- GitTagrelease name
- GitCommit git revision ID
- DefaultIndexURI
- BasePath - root directory for charmil installation
- IndexPath - path where stores the local copy of the index git repository
- InstallPath - directory for plugin installations
- BinPath - directory for the symbolic links to the installed plugin executables
