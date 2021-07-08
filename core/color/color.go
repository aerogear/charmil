// Color package is for printing a uniform set of colors for the CLI
package color

import (
	"fmt"
	"runtime"

	"github.com/fatih/color"
)

// CodeSnippet returns a colored string for code and command snippets
func CodeSnippet(format string) string {
	return color.HiMagentaString(format)
}

// Info returns a colored string for information messages
func Info(format string) string {
	return color.HiCyanString(format)
}

// Success returns a colored string for success messages
func Success(format string) string {
	return color.HiGreenString(format)
}

// Error returns a colored string for error messages
func Error(format string) string {
	return color.HiRedString(format)
}

// Bold makes a string bold
func Bold(s string) string {
	// do not bold the string if the current OS is Windows
	// Command Prompt does not support ANSI escape characters
	if runtime.GOOS == "windows" {
		return s
	}
	return fmt.Sprintf("\033[1m%v\033[0m", s)
}
