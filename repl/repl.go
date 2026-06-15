package repl

import (
  "fmt"
  "io"
  "bufio"
  "postfix/lexer"
  "postfix/parser"
)

const PROMPT = "> "

func Start(input io.Reader, output io.Writer) {
  scanner := bufio.NewScanner(input)

  for {
    fmt.Printf(PROMPT)

    if !scanner.Scan() {
      return
    }

    l := lexer.New(scanner.Text())
    p := parser.New(l)

    session := p.Parse()

    io.WriteString(output, "\t")
    io.WriteString(output, session.String())
    io.WriteString(output, "\n")
  }
}