package ast

import (
  "bytes"
)

func (infix *InfixExpression) expression() {

}

func (infix *InfixExpression) Literal() string {
  return infix.Token.Literal
}

func (infix *InfixExpression) String() string {
  var out bytes.Buffer

  out.WriteString("(")
  out.WriteString(infix.Left.String())
  out.WriteString(infix.Operator)
  out.WriteString(infix.Right.String())
  out.WriteString(")")

  return out.String()
}