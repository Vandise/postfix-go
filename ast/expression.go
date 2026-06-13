package ast

func (es *ExpressionStatement) statement() {

}

func (es *ExpressionStatement) Literal() string {
  return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
  if es.Expression != nil {
    return es.Expression.String()
  }

  return ""
}