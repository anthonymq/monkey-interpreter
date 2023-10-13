package parser

import (
	"monkey/lexer"
	"testing"
)


func TestLetStatements(t *testing.T) {
	tests := []struct {
		input              string
		expectedIdentifier string
	}{
		{"let x = 5;", "x"},
		{"let y = true;", "y"},
		{"let foobar = y;", "foobar"},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain 1 statements. got=%d",
				len(program.Statements))
		}

		stmt := program.Statements[0]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestLetStatementsErrors(t *testing.T) {
	tests := []struct {
		input               string
		expectedErrorNumber int
	}{
		{"let x 5;", 1},
		{"let = true;", 1},
		{"let y;", 1},
		{"let y;let x 5", 2},
	}

	for _, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		p.ParseProgram()
		errors := p.Errors()

		if len(errors) != tt.expectedErrorNumber {
			t.Fatalf("Bad number of error detected. got=%d expect=%d", len(errors), tt.expectedErrorNumber)
		}
	}
}
