//lexer/lexer_test.go

package lexer

import{
	"testing"
	"Interpreter/token"
}

func TestNextToke(t *testing.T){
	input := '=+(){},;'

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN , "("},
		{token.RPAREN , ")"},
		{token.LBRACE , "{"},
		{token.RBRACE , "}"},
		{token.COMMA , ","},
		{token.SEMICOLON , ";"},
		{token.EOF , ""}
	}

	l := New(input)
	for i, tt := range test{
		tok := l.NextToken()

		if tok.Type != tt.expectedType{
			t.Fatalf("test[%d] - tokentype wrong. expcted=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong. expected=%q, got=%q", i. tt.expectedLiteral, tok.Literal)
		}
	}
}