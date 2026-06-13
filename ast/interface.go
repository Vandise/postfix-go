package ast

type Node interface {
  Literal() string
}

type Statement interface {
  Node
  statement()
}

type Expression interface {
  Node
  expression()
}