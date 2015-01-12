package micropiet

func replaceTokens(in <-chan *Token, out chan<- *Token, quit <-chan struct{}, macros map[string][]*Token) {

	for {
		select {
		case tok := <-in:
			if tok.Type == TT_IDENTIFIER {
				handleIdToken(tok, macros, out)
			} else {
				out <- tok
			}
		case <-quit:
			return
		}
	}
}

func handleIdToken(tok *Token, macros map[string][]*Token, out chan<- *Token) {
	stream, ok := macros[tok.Data]
	if !ok {
		out <- tok
		return
	}
	for _, t := range stream {
		if t.Type == TT_IDENTIFIER {
			handleIdToken(t, macros, out)
		} else {
			out <- t
		}
	}
}
