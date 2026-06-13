package parser

import (
  "testing"
  "postfix/ast"
  "postfix/lexer"
)

func isValidAssignmentStatement(t *testing.T, statement ast.Statement, name string) bool {
  if statement.Literal() != "T_IDENTIFIER"  {
    t.Errorf("Statement Literal not T_IDENTIFIER, got=%q", statement.Literal())
    return false
  }

  assignmentStatement, ok := statement.(*ast.AssignmentStatement)
  if !ok {
    t.Errorf("Statement not AssignmentStatement, got=%T", statement)
    return false
  }

  if assignmentStatement.Name.Value != name {
    t.Errorf("Statement Name not %s, got=%s", name, assignmentStatement.Name.Value)
    return false
  }

  if assignmentStatement.Name.Literal() != name {
    t.Errorf("Name not %s, got=%s", name, assignmentStatement.Name.Literal())
    return false
  }

  return true
}

func Test__AssignmentStatementx(t *testing.T) {
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