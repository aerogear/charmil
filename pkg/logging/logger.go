package logging

type Logger interface {
	Output(args ...interface{})
	Error(args ...interface{})
}
