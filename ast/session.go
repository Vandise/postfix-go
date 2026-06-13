package ast

func (session *Session) Literal() string {
  if len(session.Statements) > 0 {
    return session.Statements[0].Literal()
  }

  return ""
}