package repl

import (
  "fmt"
  "io"
  "bufio"
  "postfix/token"
  "postfix/lexer"
)

const PROMPT = "> "

func Start(input io.Reader, output io.Writer) {
  scanner := bufio.NewScanner(input)

  for {
    fmt.Printf(PROMPT)

    if !scanner.Scan() {
      return
    }

    lexer := lexer.New(scanner.Text())

    for t := lexer.NextToken(); t.Type != token.T_END; t = lexer.NextToken() {
      fmt.Printf("\t%+v\n", t)
    }
  }
}