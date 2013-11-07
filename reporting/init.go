package reporting

import (
	"fmt"
	"github.com/smartystreets/goconvey/printing"
	"os"
	"strings"
	"syscall"
)

func init() {
	if !isXterm() {
		monochrome()
	}

	if isWindows() {
		success, failure, error_ = dotSuccess, dotFailure, dotError
	}
}

func BuildJsonReporter() Reporter {
	out := printing.NewPrinter(printing.NewConsole())
	return NewReporters(
		NewGoTestReporter(),
		NewJsonReporter(out))
}
func BuildDotReporter() Reporter {
	out := printing.NewPrinter(printing.NewConsole())
	return NewReporters(
		NewGoTestReporter(),
		NewDotReporter(out),
		NewProblemReporter(out),
		NewStatisticsReporter(out))
}
func BuildStoryReporter() Reporter {
	out := printing.NewPrinter(printing.NewConsole())
	return NewReporters(
		NewGoTestReporter(),
		NewStoryReporter(out),
		NewProblemReporter(out),
		NewStatisticsReporter(out))
}

var (
	newline         = "\n"
	success         = "✔"
	failure         = "✘"
	error_          = "🔥"
	skip            = "⚠"
	dotSuccess      = "."
	dotFailure      = "x"
	dotError        = "E"
	dotSkip         = "S"
	errorTemplate   = "* %s \nLine %d: - %v \n%s\n"
	failureTemplate = "* %s \nLine %d:\n%s\n"
)

var (
	greenColor  = "\033[32m"
	yellowColor = "\033[33m"
	redColor    = "\033[31m"
	resetColor  = "\033[0m"
)

// QuiteMode disables all console output symbols. This is only meant to be used
// for tests that are internal to goconvey where the output is distracting or
// otherwise not needed in the test output.
func QuietMode() {
	success, failure, error_, skip, dotSuccess, dotFailure, dotError, dotSkip = "", "", "", "", "", "", "", ""
}

func monochrome() {
	greenColor, yellowColor, redColor, resetColor = "", "", "", ""
}

func isXterm() bool {
	return strings.Contains(fmt.Sprintf("%v", os.Environ()), " TERM=isXterm")
}

// There has got to be a better way...
func isWindows() bool {
	return os.PathSeparator == '\\' && os.PathListSeparator == ';'
}
