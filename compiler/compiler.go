package compiler

import (
	"fmt"

	token "methene/ast"
	errors "methene/errors"
)

const EOL = '\r'

type Compiler struct {
	// Public
	CompiledCode string
	Error string

	// Private
	compilerBreak bool
	tokens []token.Token
	currentTokenPosition int
	currentToken token.Token
}

func NewCompiler(sourceCodeLines []token.Token) Compiler {
	return Compiler {
		// Public
		CompiledCode: "",
		Error: "",

		// Private
		compilerBreak: false,
		tokens: sourceCodeLines,
		currentTokenPosition: 0,
		currentToken: sourceCodeLines[0],
	}
}

func (this *Compiler) Compile() {
	for !this.compilerBreak {
		switch this.currentToken.TokenType {
		case token.Print:
			this.ExpectNextTokenTypeEither([]token.TokenType{token.String, token.Integer})
			this.NextToken()

			if this.currentToken.TokenType == token.String {
				this.CompiledCode += "console.log(\""+this.currentToken.TokenValue+"\");";
			}

		default:
			this.ThrowError("Unexpected token "+this.currentToken.TokenValue+" of type "+string(this.currentToken.TokenType)+" while compilation.")
		}

		if this.currentTokenPosition == len(this.tokens)-1 {
			break
		}

		this.NextToken()
	}
}

func (this *Compiler) LookNextToken() token.Token {
	return this.tokens[this.currentTokenPosition+1]
}

func (this *Compiler) LookNextTokenType() token.TokenType {
	return this.LookNextToken().TokenType
}

func (this *Compiler) NextToken() {
	if this.compilerBreak {
		return
	}

	this.currentTokenPosition += 1
	this.currentToken = this.tokens[this.currentTokenPosition]
}

func contains(s []token.TokenType, str token.TokenType) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func (this *Compiler) ExpectNextTokenTypeEither(requiredTokens []token.TokenType) {
	nextToken := this.LookNextToken()

	if !contains(requiredTokens, nextToken.TokenType) {
		this.ThrowError(fmt.Sprint("Expected either of tokens",requiredTokens,"but found '"+nextToken.TokenValue+"'."))
	}
}

func (this *Compiler) ThrowError(msg string) {
	this.Error = errors.NewError(this.currentToken.LineNumber, errors.CompilationError, msg)
	this.compilerBreak = true
}