package lexer

type Lexer struct {
	input     string
	tokens    []Token
	calculate bool
}

func NewLexer(input string) Lexer {
	return Lexer{
		input,
		[]Token{},
		true,
	}
}

func (lx *Lexer) SetInput(input string) {
	lx.input = input
	lx.calculate = true
}

func (lx *Lexer) GetTokens() []Token {
	if !lx.calculate {
		return lx.tokens
	}
	tokenizer := InitTokenizer(lx.input)
	var t Token
	for {
		t = tokenizer.NextToken()
		lx.tokens = append(lx.tokens, t)
		if t.IsEnd() {
			break
		}
	}
	lx.calculate = false
	return lx.tokens
}
