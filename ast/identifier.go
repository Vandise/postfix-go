package ast

func (identifier *Identifier) expression() {

}

func (identifier *Identifier) Literal() string {
  return identifier.Token.Literal
}

func (identifier *Identifier) String() string {
  return identifier.Value
}