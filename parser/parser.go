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

func (parser *Parser) nextToken() {
  parser.current = parser.peek
  parser.peek = parser.l.NextToken()
}

func (parser *Parser) currentTokenIs(t token.TokenType) bool {
  return parser.current.Type == t
}

func (parser *Parser) peekTokenIs(t token.TokenType) bool {
  return parser.peek.Type == t
}

func (parser *Parser) expectPeek(t token.TokenType) bool {
  if parser.peekTokenIs(t) {
    parser.nextToken()
    return true
  }

  return false
}

func New(l *lexer.Lexer) *Parser {
  parser := &Parser{ l: l }

  // set current and peek
  parser.nextToken()
  parser.nextToken()

  return parser
}

func (parser *Parser) Parse() *ast.Session {
  session := &ast.Session{}
  session.Statements = []ast.Statement{}

  for parser.current.Type != token.T_END {
    statement := parser.parseStatement()

    if statement != nil {
      session.Statements = append(session.Statements, statement)
    }

    parser.nextToken()
  }

  return session
}