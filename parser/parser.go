package parser

import (
  "postfix/token"
  "postfix/lexer"
  "postfix/ast"
)

type Parser struct {
  l *lexer.Lexer

  current token.Token
  peek    token.Token
}

func New(l *lexer.Lexer) *Parser {
  parser := &Parser{ l: l }

  // set current and peek
  parser.nextToken()
  parser.nextToken()

  return parser
}

func (parser *Parser) nextToken() {
  parser.current = parser.peek
  parser.peek = parser.l.NextToken()
}

func (parser *Parser) Parse() *ast.Session {
  return nil
}