package lexer

import "fmt"

type Lexer struct {
	input     string
	tokens    []Token
	calculate bool
	data      map[string]interface{}
}

func NewLexer(input string, data map[string]interface{}) Lexer {
	return Lexer{
		input,
		[]Token{},
		true,
		data,
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
	var prev Token
	captureSub := false
	currentSub := []Token{}
	for {
		prev = t
		t = tokenizer.NextToken()
		// { && { => start of substitution | next is token to evaluate
		if prev.Type == LSql && t.Type == LSql {
			captureSub = true
			currentSub = []Token{}
		}
		if captureSub {
			currentSub = append(currentSub, prev)
		}
		// } && } => end of substitution | next token back to path | can adjust past
		if t.Type == RSql && prev.Type == RSql {
			captureSub = false
			currentSub = append(currentSub, t)
			runSubstitutions(currentSub)
		}
		lx.tokens = append(lx.tokens, t)
		if t.IsEnd() {
			break
		}
	}
	lx.calculate = false
	return lx.tokens
}

func runSubstitutions(currentSub []Token) {
	fmt.Println(currentSub)

}
