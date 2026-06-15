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
  SUM
  PRODUCT
  PREFIX
)

var precedences = map[token.TokenType]int{
  token.T_ASSIGN: EQUALS,
  token.T_PLUS:   SUM,
  token.T_MINUS:  SUM,
  token.T_SLASH:  PRODUCT,
  token.T_STAR:   PRODUCT,
}

type Parser struct {
  l *lexer.Lexer

  current token.Token
  peek    token.Token

  prefixParseFunctions map[token.TokenType]PrefixParseFunction
  infixParseFunctions  map[token.TokenType]InfixParseFunction
}

func (parser *Parser) registerPrefixFunction(tokenType token.TokenType, handle PrefixParseFunction) {
  parser.prefixParseFunctions[tokenType] = handle
}

func (parser *Parser) registerInfixFunction(tokenType token.TokenType, handle InfixParseFunction) {
  parser.infixParseFunctions[tokenType] = handle
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

func (parser *Parser) peekPrecedence() int {
  if p, ok := precedences[parser.peek.Type]; ok {
    return p
  }

  return LOWEST
}

func (parser *Parser) currentPrecedence() int {
  if p, ok := precedences[parser.current.Type]; ok {
    return p
  }

  return LOWEST
}

func New(l *lexer.Lexer) *Parser {
  parser := &Parser{ l: l }

  parser.prefixParseFunctions = make(map[token.TokenType]PrefixParseFunction)
  parser.registerPrefixFunction(token.T_IDENTIFIER, parser.parseIdentifier)
  parser.registerPrefixFunction(token.T_INTEGER, parser.parseInteger)
  parser.registerPrefixFunction(token.T_MINUS, parser.parsePrefix)
  parser.registerPrefixFunction(token.T_OPEN_PAREN, parser.parseGroupedExpression)

  parser.infixParseFunctions = make(map[token.TokenType]InfixParseFunction)
  parser.registerInfixFunction(token.T_PLUS, parser.parseInfix)
  parser.registerInfixFunction(token.T_MINUS, parser.parseInfix)
  parser.registerInfixFunction(token.T_SLASH, parser.parseInfix)
  parser.registerInfixFunction(token.T_STAR, parser.parseInfix)


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

  left := prefix()

  for !parser.peekTokenIs(token.T_END) && precedence < parser.peekPrecedence() {
    infix := parser.infixParseFunctions[parser.peek.Type]

    if infix == nil {
      return left
    }

    parser.nextToken()

    left = infix(left)
  }

  return left
}

func (parser *Parser) parsePrefix() ast.Expression {
  prefix := &ast.PrefixExpression{
    Token:    parser.current,
    Operator: parser.current.Literal,
  }

  parser.nextToken()

  prefix.Right = parser.parseExpression(PREFIX)

  return prefix
}

func (parser *Parser) parseInfix(left ast.Expression) ast.Expression {
  infix := &ast.InfixExpression{
    Token: parser.current,
    Operator: parser.current.Literal,
    Left: left,
  }

  p := parser.currentPrecedence()

  parser.nextToken()

  infix.Right = parser.parseExpression(p)

  return infix
}

func (parser *Parser) parseGroupedExpression() ast.Expression {
  parser.nextToken()

  exp := parser.parseExpression(LOWEST)

  if !parser.expectPeek(token.T_CLOSE_PAREN) {
    return nil
  }

  return exp
}
