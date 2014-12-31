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
	for currentToken = <-tokens; currentToken.Type == TT_HASH; currentToken = <-tokens {
		name := Expect(TT_IDENTIFIER, tokens).Data
		Expect(TT_LCURL, tokens)
		macro := make([]*Token, 0)
		for t := <-tokens; t.Type != TT_RCURL; t = <-tokens {
			if t.Type == TT_LCURL || t.Type == TT_HASH {

			}
			macro = append(macro, t)
		}
		macros[name] = macro
	}
	fmt.Println("Macros done")
	for ; currentToken.Type != TT_EOF; currentToken = <-tokens {
		fmt.Printf("%s\n", currentToken.Type)
	}
	fmt.Println(macros)
	return nil, nil
}

func Expect(typ TokenType, tokens <-chan *Token) *Token {
	t := <-tokens
	if t.Type != typ {
		panic("Unexpected token")
	}
	return t
}
