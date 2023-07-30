package main

import (
	"fmt"

	tk "github.com/Hilst/app-ui/tokenizer"
)

func main() {
	fmt.Println("\nINPUT 0")
	const input_0 = "/data/wrapper/{{text}}/something"
	tokenizeInputs(input_0)
	fmt.Println("\nINPUT 1")
	const input_1 = "/data/wrapper/{{where: abc == 'def'}}/more/something"
	tokenizeInputs(input_1)
	fmt.Println("\nINPUT 2")
	const input_2 = "/data/wrapper/{{where: abc == 3.14}}/more/something"
	tokenizeInputs(input_2)
}

func tokenizeInputs(input string) {
	var lexer = tk.InitTokenizer(input)
	var t tk.Token
	for {
		t = lexer.NextToken()
		fmt.Println(t)
		if t.IsEnd() {
			break
		}
	}
}
