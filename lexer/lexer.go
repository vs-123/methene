package lexer

import (
	"fmt"

	token "methene/ast"
	errors "methene/errors"
)

const EOL = '\r'

type Lexer struct {
	// Public
	Tokens []token.Token
	Error string

	// Private
	lexerBreak bool

	sourceCodeCharacters []rune
	currentLineNumber int
	currentCharacterPosition int
	currentCharacter rune

	eatingString bool
	eatenString string
}

func NewLexer(sourceCodeLines string) Lexer {
	runeCode := []rune(sourceCodeLines)
	return Lexer {
		Tokens: []token.Token{},
		Error: "",

		lexerBreak: false,
		sourceCodeCharacters: runeCode,
		currentLineNumber: 1,
		currentCharacterPosition: 0,
		currentCharacter: runeCode[0],

		eatingString: false,
		eatenString: "",
	}
}

func (this *Lexer) Tokenize() {
	for !this.lexerBreak {
		if this.currentCharacterPosition == len(this.sourceCodeCharacters)-1 {
			this.lexerBreak = true
		}

		if !this.eatingString {
			switch this.currentCharacter {
			case '#':
				this.ExpectNextCharacterEither([]rune{'>'})
				this.NextCharacter()

				if (this.lexerBreak) {
					break
				}

				this.ExpectNextCharacterEither([]rune{'`'})
				this.NextCharacter()

				newCharacter := this.currentCharacter

				if newCharacter == '`' {
					this.Tokens = append(this.Tokens, token.Token{LineNumber: this.currentLineNumber, TokenType: token.Print, TokenValue: string(this.currentCharacter)})
					this.EatString()
				}

			default:
				this.ThrowError(fmt.Sprintf("Unexpected character '%c'.", this.currentCharacter))
			}
		} else if this.eatingString && this.currentCharacter == '`' {
			this.eatingString = false
			this.Tokens = append(this.Tokens, token.Token{LineNumber: this.currentLineNumber, TokenType: token.String, TokenValue: this.eatenString})
			this.eatenString = ""
		} else {
			this.eatenString += string(this.currentCharacter)
		}

		this.NextCharacter()
	}
}

func (this *Lexer) LookNextCharacter() rune {
	if this.currentCharacterPosition == len(this.sourceCodeCharacters) {
		return EOL
	}

	nextCharacter := this.sourceCodeCharacters[this.currentCharacterPosition + 1]
	if nextCharacter == '\n' {
		return this.LookNextCharacter()
	}

	return nextCharacter
}

func (this *Lexer) NextCharacter() {
	if this.lexerBreak {
		return
	}

	this.currentCharacterPosition += 1
	this.currentCharacter = this.sourceCodeCharacters[this.currentCharacterPosition]

	if this.currentCharacter == '\n' {
		this.currentLineNumber += 1
		this.NextCharacter()
	}
}

func contains(s []rune, str rune) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func (this *Lexer) ExpectNextCharacterEither(requiredCharacters []rune) {
	nextCharacter := this.LookNextCharacter()

	if !contains(requiredCharacters, nextCharacter) {
		this.ThrowError(fmt.Sprint("Expected either of characters ", requiredCharacters, ", found '", nextCharacter, "'"))
	}
}

func (this *Lexer) EatString() {
	this.eatingString = true
}

func (this *Lexer) ThrowError(msg string) {
	this.Error = errors.NewError(this.currentLineNumber, errors.LexicalError, msg)
	this.lexerBreak = true
}