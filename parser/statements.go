package parser

import (
  "postfix/token"
  "postfix/ast"
)

func (parser *Parser) parseStatement() ast.Statement {
  switch parser.current.Type {
    case token.T_IDENTIFIER:
      return parser.assignmentStatement()
    default:
      return nil
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

  for !parser.currentTokenIs(token.T_NEWLINE) {
    parser.nextToken()
  }

  return statement
}