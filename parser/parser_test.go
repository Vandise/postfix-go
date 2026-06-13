package parser

import (
  "testing"
  "postfix/ast"
  "postfix/lexer"
  "postfix/token"
)

func isValidAssignmentStatement(t *testing.T, statement ast.Statement, name string) bool {
  assignmentStatement, ok := statement.(*ast.AssignmentStatement)

  if !ok {
    t.Errorf("Statement not AssignmentStatement, got=%T", statement)
    return false
  }

  if assignmentStatement.Token.Type != token.T_ASSIGN {
    t.Errorf("Statement Token not %s, got=%s", token.T_IDENTIFIER, assignmentStatement.Token.Type)
    return false
  }

  if assignmentStatement.Name.Token.Type != token.T_IDENTIFIER {
    t.Errorf("Statement Name not %s, got=%s", token.T_IDENTIFIER, assignmentStatement.Name.Token.Type)
    return false
  }

  if assignmentStatement.Name.Value != name {
    t.Errorf("Name Literal not %s, got=%s", name, assignmentStatement.Name.Value)
    return false
  }

  return true
}

func Test__String(t *testing.T) {
  session := &ast.Session{
    Statements: []ast.Statement{
      &ast.AssignmentStatement{
        Token: token.Token{Type: token.T_ASSIGN, Literal: "="},
        Name: &ast.Identifier{
          Token: token.Token{Type: token.T_IDENTIFIER, Literal: "x"},
          Value: "x",
        },
        Value: &ast.Identifier{
          Token: token.Token{Type: token.T_IDENTIFIER, Literal: "y"},
          Value: "y",
        },
      },
    },
  }

  if session.String() != "x = y\n" {
    t.Fatalf("Expected x = y, got:%s", session.String())
  }
}

func Test__AssignmentStatement(t *testing.T) {
  input := `
    x = 10
    y = 5
  `

  l := lexer.New(input)
  parser := New(l)

  session := parser.Parse()

  if session == nil {
    t.Fatalf("Parse() returned nil")
  }

  if len(session.Statements) != 2 {
    t.Fatalf("Expected 2 AssignmentStatements, got=%d", len(session.Statements))
  }

  assertions := []struct {
    Expected string
  }{
    {"x"},
    {"y"},
  }

  for index, assertion := range assertions {
    statement := session.Statements[index]

    if !isValidAssignmentStatement(t, statement, assertion.Expected) {
      return
    }
  }
}