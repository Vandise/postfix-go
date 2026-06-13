package lexer

import "postfix/token"

type Lexer struct {
  input           string
  position        int
  currentPosition int
  ch              byte
}

func (lexer *Lexer) getCharacter() {
  if lexer.currentPosition >= len(lexer.input) {
    lexer.ch = token.END_TOKEN_VALUE
  } else {
    lexer.ch = lexer.input[lexer.currentPosition]
  }

  lexer.position = lexer.currentPosition
  lexer.currentPosition += 1
}

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func isDigit(ch byte) bool {
  return '0' <= ch || '9' <= ch
}

func (lexer *Lexer) readIdentifier() string {
  start := lexer.position

  for isLetter(lexer.ch) {
    lexer.getCharacter()
  }

  return lexer.input[ start : lexer.position ]
}

func (lexer *Lexer) readNumeric() string {
  start := lexer.position

  for isDigit(lexer.ch) {
    lexer.getCharacter()
  }

  return lexer.input[ start : lexer.position ]
}

func (lexer *Lexer) skipWhiteSpace() {
  for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\r' {
    lexer.getCharacter()
  }
}

func (lexer *Lexer) peek(n int) byte {
  if lexer.currentPosition + n > len(lexer.input) {
    return token.END_TOKEN_VALUE
  }

  return lexer.input[ lexer.currentPosition + n ]
}

func New(input string) *Lexer {
  lexer := &Lexer{ input: input }

  lexer.getCharacter()

  return lexer
}

func (lexer *Lexer) NextToken() token.Token {
  var t token.Token

  lexer.skipWhiteSpace()

  switch lexer.ch {
    case '+':
      t.Literal = string(lexer.ch)
      t.Type = token.T_PLUS
    case '-':
      t.Literal = string(lexer.ch)
      t.Type = token.T_MINUS
    case '*':
      t.Literal = string(lexer.ch)
      t.Type = token.T_STAR
    case '/':
      t.Literal = string(lexer.ch)
      t.Type = token.T_SLASH
    case '(':
      t.Literal = string(lexer.ch)
      t.Type = token.T_OPEN_PAREN
    case ')':
      t.Literal = string(lexer.ch)
      t.Type = token.T_CLOSE_PAREN
    case '=':
      t.Literal = string(lexer.ch)
      t.Type = token.T_ASSIGN
    case '!':
      t.Literal = string(lexer.ch)
      t.Type = token.T_BANG
    case '\n':
      t.Literal = string(lexer.ch)
      t.Type = token.T_NEWLINE
    case token.END_TOKEN_VALUE:
      t.Literal = ""
      t.Type = token.T_END
    default:
      if isLetter(lexer.ch) {
        t.Literal = lexer.readIdentifier()
        t.Type = token.T_IDENTIFIER

        return t
      }

      if isDigit(lexer.ch) {
        t.Literal = lexer.readNumeric()
        t.Type = token.T_INTEGER

        return t
      }
  }

  lexer.getCharacter()

  return t
}