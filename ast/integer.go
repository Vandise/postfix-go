package ast

func (integer *IntegerLiteral) expression() {

}

func (integer *IntegerLiteral) Literal() string {
  return integer.Token.Literal
}

func (integer *IntegerLiteral) String() string {
  return integer.Token.Literal
}