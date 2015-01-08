package micropiet

type parser struct {
	toks <-chan *Token
}

func (p *parser) next() *Token {
	return <-p.toks
}

func (p *parser) expect(typ TokenType) *Token {
	t := p.next()
	if t.Type != typ {
		panic("Unexpected token")
	}
	return t
}

func Parse(text string) (*Program, error) {
	tokens := make(chan *Token)
	go Tokenize(text, tokens)
	p := parser{tokens}
	macros := make(map[string][]*Token)
	var currentToken *Token
	//Parse macros
	for currentToken = p.next(); currentToken.Type == TT_HASH; currentToken = p.next() {
		name := p.expect(TT_IDENTIFIER).Data
		p.expect(TT_LCURL)
		macro := []*Token{}
		for t := p.next(); t.Type != TT_RCURL; t = p.next() {
			if t.Type == TT_LCURL || t.Type == TT_HASH {
				panic("no nested macros")
			}
			macro = append(macro, t)
		}
		macros[name] = macro
	}
	prog := NewProgram()
	for ; currentToken.Type != TT_EOF; currentToken = <-tokens {
		switch currentToken.Type {
		case TT_INTEGER:
			prog.AddCommand(NewPushCommand(currentToken.Data))
		case TT_ADD:
			fallthrough
		case TT_MUL:
			fallthrough
		case TT_DIV:
			fallthrough
		case TT_MOD:
			fallthrough
		case TT_GT:
			fallthrough
		case TT_SUB:
			prog.AddCommand(NewArithmeticCommand(currentToken.Data))
		case TT_POP:
			prog.AddCommand(NewPopCommand())
		default:
			panic("Unexpected token received: " + currentToken.Type.String())
		}
	}
	return prog, nil
}
