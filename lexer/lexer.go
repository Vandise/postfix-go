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
    lexer.ch = 0
  } else {
    lexer.ch = lexer.input[lexer.currentPosition]
  }

  lexer.position = lexer.currentPosition
  lexer.currentPosition += 1
}

func isLetter(ch byte) bool {
  return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z'
}

func (lexer *Lexer) readIdentifier() string {
  start := lexer.position

  for isLetter(lexer.ch) {
    lexer.getCharacter()
  }

  return lexer.input[ start : lexer.position ]
}

func New(input string) *Lexer {
  lexer := &Lexer{ input: input }

  lexer.getCharacter()

  return lexer
}

func (lexer *Lexer) NextToken() token.Token {
  var t token.Token

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
    case 0:
      t.Literal = ""
      t.Type = token.T_END
    default:
      if isLetter(lexer.ch) {
        t.Literal = lexer.readIdentifier()
        t.Type = token.T_IDENTIFIER

        return t
      }
  }

  lexer.getCharacter()

  return t
}