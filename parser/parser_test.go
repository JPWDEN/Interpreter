package parser

import (
	"testing"

	"github.com/Interpreter/lexer"
)

func testLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 838383;`
	lex := lexer.NewLexer(input)
	parse := NewParser(lex)
	program := parse.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}
}

func testLetStatement(t *testing.T) {

}
