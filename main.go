package main

import (
	"fmt"

	lx "github.com/Hilst/app-ui/lexer"
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
	lexer := lx.NewLexer(input, make(map[string]interface{}))
	tokens := lexer.GetTokens()
	for i := 0; i < len(tokens); i++ {
		fmt.Printf("%d => %s\n", i, tokens[i])
	}
}
