package main

import (
  "os"
  "postfix/repl"
)

func main() {
  repl.Start(os.Stdin, os.Stdout)
}