package token

type TokenType string

type Token struct {
  Type    TokenType
  Literal string
}

const (
  T_ILLEGAL     = "T_ILLEGAL"
  T_IDENTIFIER  = "T_IDENTIFIER"
  T_ASSIGN      = "T_ASSIGN"
  T_PLUS        = "T_PLUS"
  T_MINUS       = "T_MINUS"
  T_SLASH       = "T_SLASH"
  T_STAR        = "T_STAR"
  T_OPEN_PAREN  = "T_OPEN_PAREN"
  T_CLOSE_PAREN = "T_CLOSE_PAREN"
)