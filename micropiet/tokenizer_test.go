package micropiet

import (
	"testing"
)

func checkToken(t *testing.T, actual TokenType, expected TokenType) {
	if actual != expected {
		t.Errorf("Expected %s token, but found %s.", actual, expected)
	}
}

func TestEmptyString(t *testing.T) {
	out := make(chan *Token)
	go Tokenize("", out)
	tok := <-out
	checkToken(t, tok.Type, TT_EOF)
}

func TestSingleInteger(t *testing.T) {
	out := make(chan *Token)
	go Tokenize("   \t123", out)
	tok := <-out
	checkToken(t, tok.Type, TT_INTEGER)
	if tok.Data != "123" {
		t.Errorf("Bad data. %s not %s", tok.Data, "123")
	}
	tok = <-out
	checkToken(t, tok.Type, TT_EOF)
}

func TestTwoIntegers(t *testing.T) {
	out := make(chan *Token)
	go Tokenize("123 456", out)
	tok := <-out
	checkToken(t, tok.Type, TT_INTEGER)
	if tok.Data != "123" {
		t.Errorf("Bad data. %s not %s", tok.Data, "123")
	}
	tok = <-out
	checkToken(t, tok.Type, TT_INTEGER)
	if tok.Data != "456" {
		t.Errorf("Bad data. %s not %s", tok.Data, "456")
	}
}

func TestComment(t *testing.T) {
	out := make(chan *Token)
	go Tokenize("123 //567 \n456", out)
	tok := <-out
	checkToken(t, tok.Type, TT_INTEGER)
	if tok.Data != "123" {
		t.Errorf("Bad data. %s not %s", tok.Data, "123")
	}
	tok = <-out
	checkToken(t, tok.Type, TT_INTEGER)
	if tok.Data != "456" {
		t.Errorf("Bad data. %s not %s", tok.Data, "456")
	}
}

func TestParse(t *testing.T) {
	Parse(`
	#moveUp{ //move nth cell to top. tip of stack is 0
	1 2 - r}
	#moveAfter{ //move to after the nth cell
	1 r}
	//read width and height, calculate size
	NdNd 3 moveUp`)
}
