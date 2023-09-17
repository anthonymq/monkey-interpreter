package parser

import (
	"monkey/ast"
	"monkey/lexer"
	"monkey/token"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		l *lexer.Lexer
	}
	lexer := lexer.New("let x = 5;")
	tests := []struct {
		name string
		args args
		want *Parser
	}{
		{"constructor", args{l: lexer}, &Parser{l: lexer, curToken: token.Token{Type: token.LET, Literal: "let"}, peekToken: token.Token{Type: token.IDENT, Literal: "x"}, errors: []string{}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

// func TestParser_ParseProgram(t *testing.T) {
// 	tests := []struct {
// 		name  string
// 		input *lexer.Lexer
// 		want  *ast.Program
// 	}{
// 		{
// 			"LetStatement",
// 			lexer.New("let x = 5;"),
// 			&ast.Program{
// 				Statements: []ast.Statement{
// 					&ast.LetStatement{
// 						Token: token.Token{Type: token.LET, Literal: "let"},
// 						Name: &ast.Identifier{
// 							Token: token.Token{Type: token.IDENT, Literal: "x"},
// 							Value: "x",
// 						},
// 						Value: &ast.Identifier{
// 							Token: token.Token{Type: token.IDENT, Literal: "5"},
// 							Value: "5",
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			"LetStatement2",
// 			lexer.New("let x = 5;let y = 6;"),
// 			&ast.Program{
// 				Statements: []ast.Statement{
// 					&ast.LetStatement{
// 						Token: token.Token{Type: token.LET, Literal: "let"},
// 						Name: &ast.Identifier{
// 							Token: token.Token{Type: token.IDENT, Literal: "x"},
// 							Value: "x",
// 						},
// 						Value: &ast.Identifier{
// 							Token: token.Token{Type: token.IDENT, Literal: "5"},
// 							Value: "5",
// 						},
// 					},
// 					&ast.LetStatement{
// 						Token: token.Token{Type: token.LET, Literal: "let"},
// 						Name: &ast.Identifier{
// 							Token: token.Token{Type: token.IDENT, Literal: "y"},
// 							Value: "y",
// 						},
// 						Value: &ast.Identifier{
// 							Token: token.Token{Type: token.IDENT, Literal: "6"},
// 							Value: "6",
// 						},
// 					},
// 				},
// 			},
// 		},
// 		{
// 			"LetStatementError",
// 			lexer.New("let = 5;"),
// 			&ast.Program{
// 				Statements: []ast.Statement{
// 					&ast.LetStatement{
// 						Token: token.Token{Type: token.LET, Literal: "let"},
// 						Name: &ast.Identifier{
// 							Token: token.Token{Type: token.IDENT, Literal: "x"},
// 							Value: "x",
// 						},
// 						Value: &ast.Identifier{
// 							Token: token.Token{Type: token.IDENT, Literal: "5"},
// 							Value: "5",
// 						},
// 					},
// 				},
// 			},
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			p := New(tt.input)
// 			program := p.ParseProgram()

// 			checkParserErrors(t, p)

// 			if program == nil {
// 				t.Fatalf("ParseProgram() returned nil")
// 			}

// 			if len(program.Statements) != len(tt.want.Statements) {
// 				t.Fatalf("program.Statements does not contain %d statements. got=%d", len(tt.want.Statements), len(program.Statements))
// 			}
// 			for i, statement := range program.Statements {
// 				expectedLetStatement := tt.want.Statements[i].(*ast.LetStatement)

// 				if statement.TokenLiteral() != tt.want.TokenLiteral() {
// 					t.Errorf("statement.TokenLiteral not '%q'. got=%q", tt.want.TokenLiteral(), statement.TokenLiteral())
// 				}

// 				letStmt, ok := statement.(*ast.LetStatement)
// 				if !ok {
// 					t.Errorf("s not *ast.LetStatement. got=%T", statement)
// 				}

// 				expectedName := expectedLetStatement.Name
// 				if letStmt.Name.Value != expectedName.Value {
// 					t.Errorf("letStmt.Name.Value not '%s'. got=%s", expectedName.Value, letStmt.Name.Value)
// 				}

// 				if letStmt.Name.TokenLiteral() != expectedName.TokenLiteral() {
// 					t.Errorf("s.Name not '%s'. got=%s", expectedName.TokenLiteral(), letStmt.Name.TokenLiteral())
// 				}
// 			}

// 		})
// 	}
// }

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
		input              string
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

		if len(errors) !=  tt.expectedErrorNumber {
			t.Fatalf("Bad number of error detected. got=%d expect=%d", len(errors), tt.expectedErrorNumber)
		}
		// checkParserErrors(t, p)

		// if len(program.Statements) != 1 {
		// 	t.Fatalf("program.Statements does not contain 1 statements. got=%d",
		// 		len(program.Statements))
		// }

		// stmt := program.Statements[0]
		// if !testLetStatement(t, stmt, tt.expectedIdentifier) {
		// 	return
		// }
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s'. got=%s", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("s.Name not '%s'. got=%s", name, letStmt.Name)
		return false
	}

	return true
}

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}
	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}
