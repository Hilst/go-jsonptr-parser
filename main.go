package main

import (
	"fmt"

	lx "github.com/Hilst/app-ui/lexer"
	stx "github.com/Hilst/app-ui/syntax"
)

func main() {
	fmt.Println("\nINPUT 0")
	const input_0 = "/data/wrapper/{{where: abc == 'def'}}/more/something"
	tokenizeInputs(input_0)
	fmt.Println("\nINPUT 1")
	const input_1 = "/data/wrapper/{{where: abc == 3.14}}/more/{{index}}/bit_more/something"
	tokenizeInputs(input_1)
	fmt.Println("\nINPUT 2")
	const input_2 = "/data/wrapper/{{index}}/more/{{where: abc == 3.14}}/bit_more/something"
	tokenizeInputs(input_2)
}

func tokenizeInputs(input string) {
	lexer := lx.NewLexer(input)
	tokens := lexer.GetTokens()
	for i := 0; i < len(tokens); i++ {
		fmt.Printf("%d => %s\n", i, tokens[i])
	}
	syntaxAnalyser := stx.NewSyntaxAnalyser(tokens)
	subs := syntaxAnalyser.GetSubstitutions(stx.WhereSubs)
	for i := 0; i < len(subs); i++ {
		fmt.Println(lx.Where)
		fmt.Println(subs)
	}
	subs = syntaxAnalyser.GetSubstitutions(stx.IndexSubs)
	for i := 0; i < len(subs); i++ {
		fmt.Println(lx.Index)
		fmt.Println(subs)
	}
}
