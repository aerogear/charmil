package validator

var (
	LengthRule      = "LENGTH_RULE"
	MustPresentRule = "MUST_PRESENT_RULE"
)

type stats struct {
	num        int
	num_failed int
	errors     []Error
}
