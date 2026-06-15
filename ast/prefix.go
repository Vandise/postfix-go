package ast

import (
  "bytes"
)

func (prefix *PrefixExpression) expression() {

}

func (prefix *PrefixExpression) Literal() string {
  return prefix.Token.Literal
}

func (prefix *PrefixExpression) String() string {
  var out bytes.Buffer

  out.WriteString(prefix.Operator)
  out.WriteString(prefix.Right.String())

  return out.String()
}