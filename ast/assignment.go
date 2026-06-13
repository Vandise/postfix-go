package ast

//import "postfix/token"

func (assignment *AssignmentStatement) statement() {

}

func (assignment *AssignmentStatement) Literal() string {
  return assignment.Token.Literal
}