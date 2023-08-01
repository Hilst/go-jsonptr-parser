package lexer

import "fmt"

type TokenType string

const (
	Illegal TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	LSql TokenType = "{"
	RSql TokenType = "}"

	Where   TokenType = "WHERE"
	Equal   TokenType = "="
	Greater TokenType = ">"
	Less    TokenType = "<"

	Number TokenType = "NUMBER"

	Ident  TokenType = "IDENT"
	SQuote TokenType = "'"

	Index TokenType = "INDEX"
)

type Token struct {
	Type    TokenType
	Literal string
}

func CreateToken(Type TokenType, literal string) Token {
	return Token{
		Type,
		literal,
	}
}
func (t *Token) IsEnd() bool {
	return t.Type == Illegal || t.Type == EOF
}
func (t Token) Format(f fmt.State, c rune) {
	var output = "< "
	if t.Type == Ident || t.Type == Number {
		output += string(t.Type)
		output += ": "
		output += t.Literal
	} else {
		output += string(t.Type)
	}
	output += " >"
	f.Write([]byte(output))
}
