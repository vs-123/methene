package token

type TokenType int

const (
	Beginning TokenType = iota
	// Data types
    String
    Integer
	// Commands
    Print
	// EOL
    EOL
)

type Token struct {
	LineNumber int
	TokenType TokenType
	TokenValue string
}