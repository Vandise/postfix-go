package parser

import (
  "strconv"
  "postfix/token"
  "postfix/lexer"
  "postfix/ast"
)

//
// Pratt-style Parser, mapping Tokens to Function Definitions
//
type (
  PrefixParseFunction func() ast.Expression
  InfixParseFunction  func(ast.Expression) ast.Expression
)

const (
  _ int = iota
  LOWEST
  EQUALS
  LTGT
  SUM
  PRODUCT
  PREFIX
)

type Parser struct {
  l *lexer.Lexer

  current token.Token
  peek    token.Token

  prefixParseFunctions map[token.TokenType]PrefixParseFunction
  infixParseFunction   map[token.TokenType]InfixParseFunction
}

func (parser *Parser) registerPrefixFunction(tokenType token.TokenType, handle PrefixParseFunction) {
  parser.prefixParseFunctions[tokenType] = handle
}

func (parser *Parser) registerInfixFunction(tokenType token.TokenType, handle InfixParseFunction) {
  parser.infixParseFunction[tokenType] = handle
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

  parser.prefixParseFunctions = make(map[token.TokenType]PrefixParseFunction)
  parser.registerPrefixFunction(token.T_IDENTIFIER, parser.parseIdentifier)
  parser.registerPrefixFunction(token.T_INTEGER, parser.parseInteger)

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

func (parser *Parser) parseIdentifier() ast.Expression {
  return &ast.Identifier{ Token: parser.current, Value: parser.current.Literal }
}

func (parser *Parser) parseInteger() ast.Expression {
  literal := &ast.IntegerLiteral{ Token: parser.current }

  value, err := strconv.ParseInt(parser.current.Literal, 0, 64)

  if err != nil {
    return nil
  }

  literal.Value = value

  return literal
}

func (parser *Parser) parseExpression(precedence int) ast.Expression {
  prefix := parser.prefixParseFunctions[parser.current.Type]

  if prefix == nil {
    return nil
  }

  return prefix()
}