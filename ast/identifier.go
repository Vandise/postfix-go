package ast

//import "postfix/token"

func (identifier *Identifier) expression() {

}

func (identifier *Identifier) Literal() string {
  return identifier.Token.Literal
}