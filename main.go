package main

import (
	"fmt"
	"os"
	"strings"

	// token "methene/ast"
	compiler "methene/compiler"
	lexer "methene/lexer"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func UNUSED(x ...interface{}) {}

func main() {
	sourceCode, err := os.ReadFile("test.mt")
	check(err)

	sourceCodeString := string(sourceCode)

	cleanSourceCode := strings.ReplaceAll(sourceCodeString, "\r", "")

	fmt.Println(sourceCodeString)

	lexerStruct := lexer.NewLexer(cleanSourceCode)
	lexerStruct.Tokenize()
	fmt.Println("Tokens: ", lexerStruct.Tokens)

	compilerStruct := compiler.NewCompiler(lexerStruct.Tokens)
	compilerStruct.Compile()
	fmt.Println("Compiled code:\n" + compilerStruct.CompiledCode)
}
