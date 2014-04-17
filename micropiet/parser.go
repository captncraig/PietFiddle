package micropiet

import (
	"fmt"
)

type Program struct {
	Macros   map[string][]*Token
	Commands []*Command
}
type Command interface{}

func Parse(text string) (*Program, error) {
	tokens := make(chan *Token)
	go Tokenize(text, tokens)
	macros := make(map[string][]*Token)
	var currentToken *Token
	//Parse macros
	for currentToken = <-tokens; currentToken.Type == "Hash"; currentToken = <-tokens {
		name := Expect("Identifier", tokens).Data
		Expect("LCurl", tokens)
		macro := make([]*Token, 0)
		for t := <-tokens; t.Type != "RCurl"; t = <-tokens {
			macro = append(macro, t)
		}
		macros[name] = macro
	}
	fmt.Println("MAcros done")
	for ; currentToken.Type != "EOF"; currentToken = <-tokens {
		fmt.Printf("B%s\n", currentToken.Type)
	}
	fmt.Println(macros)
	return nil, nil
}

func Expect(typ string, tokens <-chan *Token) *Token {
	t := <-tokens
	if t.Type != typ {
		panic("Unexpected token")
	}
	return t
}
