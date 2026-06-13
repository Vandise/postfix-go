package ast

import "postfix/token"

type Session struct {
  Statements []Statement
}

type Identifier struct {
  Token token.Token
  Value string
}

type AssignmentStatement struct {
  Token  token.Token
  Name   *Identifier
  Value  Expression
}