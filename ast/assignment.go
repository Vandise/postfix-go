package ast

import (
  "bytes"
)

func (assignment *AssignmentStatement) statement() {

}

func (assignment *AssignmentStatement) Literal() string {
  return assignment.Token.Literal
}

func (assignment *AssignmentStatement) String() string {
  var out bytes.Buffer

  out.WriteString(assignment.Name.String())
  out.WriteString(" = ")

  if assignment.Value != nil {
    out.WriteString(assignment.Value.String())
  }

  out.WriteString("\n")

  return out.String()
}