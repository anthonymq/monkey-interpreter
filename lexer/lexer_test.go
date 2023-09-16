package lexer

import (
	"monkey/token"
	"testing"
)

func assertToken(t *testing.T, testNb int, currentToken token.Token, expectedToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}) {
	if currentToken.Type != expectedToken.expectedType {
		t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
			testNb, expectedToken.expectedType, currentToken.Type)
	}
	if currentToken.Literal != expectedToken.expectedLiteral {
		t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
			testNb, expectedToken.expectedLiteral, currentToken.Literal)
	}
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		assertToken(t, i, tok, tt)
	}
}

func TestNextTokenExtended(t *testing.T) {
	input := `let five = 5;
let ten = 10;
   let add = fn(x, y) {
     x + y;
};
   let result = add(five, ten);
   `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		assertToken(t, i, tok, tt)
	}

}

func TestNewOperators(t *testing.T) {
	input := `
	!-/*5;
	5 < 10 > 5;
   `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},

		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},

		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		assertToken(t, i, tok, tt)
	}
}

func TestKeywordTrue(t *testing.T) {
	input := `
	true;
   `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		assertToken(t, i, tok, tt)
	}
}

func TestKeywordFalse(t *testing.T) {
	input := `
	false;
   `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		assertToken(t, i, tok, tt)
	}
}

func TestKeywordIf(t *testing.T) {
	input := `
	if;
   `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.IF, "if"},
		{token.SEMICOLON, ";"},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		assertToken(t, i, tok, tt)
	}
}

func TestKeywordElse(t *testing.T) {
	input := `
	else;
   `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ELSE, "else"},
		{token.SEMICOLON, ";"},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		assertToken(t, i, tok, tt)
	}
}

func TestKeywordReturn(t *testing.T) {
	input := `
	return;
   `
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.RETURN, "return"},
		{token.SEMICOLON, ";"},
	}
	l := New(input)
	for i, tt := range tests {
		tok := l.NextToken()
		assertToken(t, i, tok, tt)
	}
}
