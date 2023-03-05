package errors

import "fmt"

type ErrorType string

const (
	LexicalError ErrorType = "Lexical Error"
	CompilationError = "Compilation Error"
)

func NewError(lineNumber int, errorType ErrorType, msg string) string {
	return fmt.Sprintf("%s at line %d: %s", errorType, lineNumber, msg)
}