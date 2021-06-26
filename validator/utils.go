package validator

// Define rules as string
var (
	LengthRule      = "LENGTH_RULE"
	MustPresentRule = "MUST_PRESENT_RULE"
)

// stats is used for providing info
// about validation of commands
type stats struct {
	// number of commands checked
	num int
	// number of commands having errors
	num_failed int
	// errors in command
	errors []Error
}
