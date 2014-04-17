package micropiet

import (
	//"fmt"
	"regexp"
	"strings"
)

type Token struct {
	Type string
	Data string
	line int
	col  int
}

type tokenType struct {
	Regex *regexp.Regexp
	Name  string
}

var tokens []tokenType = buildTokenMap()

func buildTokenMap() []tokenType {
	return []tokenType{
		//Possible tokens in order of importance
		tokenType{regexp.MustCompile("^\\s+"), "Whitespace"},
		tokenType{regexp.MustCompile("^//.*$"), "Comment"},
		tokenType{regexp.MustCompile("^[0-9]+"), "Integer"},
		tokenType{regexp.MustCompile("^p"), "Pop"},
		tokenType{regexp.MustCompile("^\\+"), "Add"},
		tokenType{regexp.MustCompile("^-"), "Sub"},
		tokenType{regexp.MustCompile("^/"), "Div"},
		tokenType{regexp.MustCompile("^\\*"), "Mul"},
		tokenType{regexp.MustCompile("^%"), "Mod"},
		tokenType{regexp.MustCompile("^!"), "Not"},
		tokenType{regexp.MustCompile("^>"), "GT"},
		tokenType{regexp.MustCompile("^d"), "Dup"},
		tokenType{regexp.MustCompile("^r"), "Roll"},
		tokenType{regexp.MustCompile("^n"), "InC"},
		tokenType{regexp.MustCompile("^N"), "InN"},
		tokenType{regexp.MustCompile("^o"), "OutC"},
		tokenType{regexp.MustCompile("^O"), "OutN"},
		tokenType{regexp.MustCompile("^jmp"), "Jmp"},
		tokenType{regexp.MustCompile("^br"), "Br"},
		tokenType{regexp.MustCompile("^#"), "Hash"},
		tokenType{regexp.MustCompile("^:"), "Colon"},
		tokenType{regexp.MustCompile("^{"), "LCurl"},
		tokenType{regexp.MustCompile("^}"), "RCurl"},
		//Identifier last
		tokenType{regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9]*"), "Identifier"},
	}
}

func Tokenize(input string, output chan<- *Token) {
	lines := strings.Split(input, "\n")
	for i, line := range lines {
		tokenizeLine(i, line, output)
	}
	output <- &Token{"EOF", "", 0, 0}
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
				if isSignifigant(tok.Name) {
					output <- &Token{tok.Name, match, i, ch}
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

func isSignifigant(tokenName string) bool {
	return tokenName != "Whitespace" && tokenName != "Comment"
}
