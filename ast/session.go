package ast

import (
  "bytes"
)

func (session *Session) Literal() string {
  if len(session.Statements) > 0 {
    return session.Statements[0].Literal()
  }

  return ""
}

func (session *Session) String() string {
  var out bytes.Buffer

  for _, statement := range session.Statements {
    out.WriteString(statement.String())
  }

  return out.String()
}