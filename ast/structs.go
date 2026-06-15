package ast

import "postfix/token"

type Session struct {
  Statements []Statement
}

type Identifier struct {
  Token token.Token
  Value string
}

type IntegerLiteral struct {
  Token token.Token
  Value int64
}

type AssignmentStatement struct {
  Token  token.Token
  Name   *Identifier
  Value  Expression
}

// expressions can also be a statement
type ExpressionStatement struct {
  Token token.Token
  Expression  Expression
}