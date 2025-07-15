package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/richardktran/monkey-language/lexer"
	"github.com/richardktran/monkey-language/token"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "" {
			continue
		}

		if line == "exit" {
			return
		}

		l := lexer.NewLexer(line)
		for {
			tok := l.NextToken()

			if tok.Type == token.EOF {
				return
			}
			fmt.Printf("%+v\n", tok)
		}
	}
}
