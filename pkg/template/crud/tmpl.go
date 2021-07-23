package crud

import _ "embed"

var (
	// createTemplate stores the template of the `create.go` file as a string
	//go:embed "create.go"
	createTemplate string

	// deleteTemplate stores the template of the `delete.go` file as a string
	//go:embed "delete.go"
	deleteTemplate string

	// describeTemplate stores the template of the `describe.go` file as a string
	//go:embed "describe.go"
	describeTemplate string

	// listTemplate stores the template of the `list.go` file as a string
	//go:embed "list.go"
	listTemplate string

	// useTemplate stores the template of the `use.go` file as a string
	//go:embed "use.go"
	useTemplate string

	// Maps the template names to their template strings
	TmplMap = map[string]string{
		"create":   createTemplate,
		"delete":   deleteTemplate,
		"describe": describeTemplate,
		"list":     listTemplate,
		"use":      useTemplate,
	}
)
