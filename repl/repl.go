package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/Interpreter/lexer"
	"github.com/Interpreter/token"
)

//PROMPT explanation becase in some ways, go is ridiculous
const PROMPT = ">>"

//Start the command line prompt, read in a line of text, and produce tokens
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := lexer.New(line)
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
