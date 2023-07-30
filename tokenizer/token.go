package tokenizer

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
)

type Token struct {
	type_   TokenType
	literal string
}

func CreateToken(type_ TokenType, literal string) Token {
	return Token{
		type_,
		literal,
	}
}
func (t *Token) IsEnd() bool {
	return t.type_ == Illegal || t.type_ == EOF
}
func (t Token) Format(f fmt.State, c rune) {
	var output = "< "
	if t.type_ == Ident || t.type_ == Number {
		output += string(t.type_)
		output += ": "
		output += t.literal
	} else {
		output += string(t.type_)
	}
	output += " >"
	f.Write([]byte(output))
}
