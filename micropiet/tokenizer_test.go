package micropiet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyString(t *testing.T) {
	out := make(chan *Token)
	go Tokenize("", out)
	tok := <-out
	assert.Equal(t, tok.Type, "EOF")
}

func TestSingleInteger(t *testing.T) {
	out := make(chan *Token)
	go Tokenize("   \t123", out)
	tok := <-out
	assert.Equal(t, tok.Type, "Integer")
	assert.Equal(t, tok.Data, "123")
	tok = <-out
	assert.Equal(t, tok.Type, "EOF")
}

func TestTwoIntegers(t *testing.T) {
	out := make(chan *Token)
	go Tokenize("123 456", out)
	tok := <-out
	assert.Equal(t, tok.Type, "Integer")
	assert.Equal(t, tok.Data, "123")
	tok = <-out
	assert.Equal(t, tok.Type, "Integer")
	assert.Equal(t, tok.Data, "456")
}

func TestComment(t *testing.T) {
	out := make(chan *Token)
	go Tokenize("123 //567 \n456", out)
	tok := <-out
	assert.Equal(t, tok.Type, "Integer")
	assert.Equal(t, tok.Data, "123")
	tok = <-out
	assert.Equal(t, tok.Type, "Integer")
	assert.Equal(t, tok.Data, "456")
}

func TestParse(t *testing.T) {
	Parse("#moveUp{ //move nth cell to top. tip of stack is 0\n1 2 - r}#moveAfter{ //move to after the nth cell\n1 r}//read width and height, calculate size\nNdNd 3 moveUp ")
}
