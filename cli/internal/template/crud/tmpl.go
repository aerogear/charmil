package crud

import "embed"

// CrudTemplates stores embedded contents of all the CRUD template files
//go:embed *
var CrudTemplates embed.FS
