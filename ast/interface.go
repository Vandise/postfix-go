package ast

type Node interface {
  Literal() string
  String()  string
}

type Statement interface {
  Node
  statement()
}

type Expression interface {
  Node
  expression()
}