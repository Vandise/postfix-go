package parser

import (
  "postfix/token"
  "postfix/ast"
)

func (parser *Parser) parseStatement() ast.Statement {
  switch parser.current.Type {
    case token.T_IDENTIFIER:
      if parser.peekTokenIs(token.T_ASSIGN) {
        return parser.assignmentStatement()
      }
      fallthrough
    default:
      return parser.expressionStatement()
  }
}

func (parser *Parser) assignmentStatement() *ast.AssignmentStatement {
  idType := parser.current
  id := parser.current.Literal

  if !parser.expectPeek(token.T_ASSIGN) {
    return nil
  }

  statement := &ast.AssignmentStatement{ Token: parser.current }

  statement.Name = &ast.Identifier{ Token: idType, Value: id }

  parser.nextToken()

  return statement
}

func (parser *Parser) expressionStatement() *ast.ExpressionStatement {
  statement := &ast.ExpressionStatement{ Token: parser.current }
  statement.Expression = parser.parseExpression(LOWEST)

  parser.nextToken()

  return statement
}