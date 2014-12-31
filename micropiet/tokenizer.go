package micropiet

import (
	"fmt"
	"regexp"
	"strings"
)

type Token struct {
	Type TokenType
	Data string
	line int
	col  int
}

func (t *Token) String() string {
	return fmt.Sprintf("%s(%s)", t.Type, t.Data)
}

//go:generate stringer -type=TokenType
type TokenType byte

const (
	TT_WHITESPACE TokenType = iota
	TT_COMMENT
	TT_INTEGER
	TT_POP
	TT_ADD
	TT_SUB
	TT_DIV
	TT_MUL
	TT_MOD
	TT_NOT
	TT_GT
	TT_DUP
	TT_ROLL
	TT_INPUT
	TT_INPUT_NUM
	TT_OUTPUT
	TT_OUTPUT_NUM
	TT_JMP
	TT_PRINT
	TT_BRANCH
	TT_SWITCH
	TT_IDENTIFIER
	TT_LABEL
	TT_HASH
	TT_LCURL
	TT_RCURL
	TT_STRING
	TT_EOF
)

type tokenDef struct {
	Regex *regexp.Regexp
	TT    TokenType
}

var tokens []tokenDef = buildTokenMap()

func buildTokenMap() []tokenDef {
	return []tokenDef{
		//Possible tokens in order of importance
		tokenDef{regexp.MustCompile("^\\s+"), TT_WHITESPACE},
		tokenDef{regexp.MustCompile("^//.*$"), TT_COMMENT},
		tokenDef{regexp.MustCompile("^[0-9]+"), TT_INTEGER},
		tokenDef{regexp.MustCompile("^p"), TT_POP},
		tokenDef{regexp.MustCompile("^\\+"), TT_ADD},
		tokenDef{regexp.MustCompile("^-"), TT_SUB},
		tokenDef{regexp.MustCompile("^/"), TT_DIV},
		tokenDef{regexp.MustCompile("^\\*"), TT_MUL},
		tokenDef{regexp.MustCompile("^%"), TT_MOD},
		tokenDef{regexp.MustCompile("^!"), TT_NOT},
		tokenDef{regexp.MustCompile("^>"), TT_GT},
		tokenDef{regexp.MustCompile("^d"), TT_DUP},
		tokenDef{regexp.MustCompile("^r"), TT_ROLL},
		tokenDef{regexp.MustCompile("^n"), TT_INPUT},
		tokenDef{regexp.MustCompile("^N"), TT_INPUT_NUM},
		tokenDef{regexp.MustCompile("^o"), TT_OUTPUT},
		tokenDef{regexp.MustCompile("^O"), TT_OUTPUT_NUM},
		tokenDef{regexp.MustCompile("^jmp"), TT_JMP},
		tokenDef{regexp.MustCompile("^print"), TT_PRINT},
		tokenDef{regexp.MustCompile("^br"), TT_BRANCH},
		tokenDef{regexp.MustCompile("^#"), TT_HASH},
		tokenDef{regexp.MustCompile("^:[a-zA-Z][a-zA-Z0-9]*"), TT_LABEL},
		tokenDef{regexp.MustCompile("^{"), TT_LCURL},
		tokenDef{regexp.MustCompile("^}"), TT_RCURL},
		//Identifier last
		tokenDef{regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]*"), TT_IDENTIFIER},
	}
}

func Tokenize(input string, output chan<- *Token) {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		tokenizeLine(i, line, output)
	}
	output <- &Token{TT_EOF, "", 0, 0}
}

func tokenizeLine(i int, line string, output chan<- *Token) {
	ch := 0

	for {
		if len(line) <= 0 {
			return
		}
		found := false
		for _, tok := range tokens {
			match := tok.Regex.FindString(line)
			if match != "" {
				found = true
				if isSignifigant(tok.TT) {
					output <- &Token{tok.TT, match, i, ch}
				}
				ch += len(match)
				line = line[len(match):]
				break
			}
		}
		if !found {
			panic("No Good Tokens!?!?!")
		}
	}
}

func isSignifigant(tt TokenType) bool {
	return tt != TT_WHITESPACE && tt != TT_COMMENT
}
