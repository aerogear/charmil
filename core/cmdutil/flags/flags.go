// flags package is a helper package for processing and interactive command line flags
package flags

var (
	ValidOutputFormats       = []string{"json", "yml", "yaml"}
	CredentialsOutputFormats = []string{"env", "json", "properties"}
)

// IsValidInput checks if the input value is in the range of valid values
func IsValidInput(input string, validValues ...string) bool {
	for _, b := range validValues {
		if input == b {
			return true
		}
	}

	return false
}
