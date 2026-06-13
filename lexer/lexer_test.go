package lexer

import(
  "testing"
  "postfix/token"
)

func Test__Tokens(t *testing.T) {
  assertions := []struct {
    ExpectedType    token.TokenType
    ExpectedLiteral string
  }{
    { token.T_PLUS,   "+" },
    { token.T_MINUS,  "-" },
    { token.T_STAR,   "*" },
    { token.T_SLASH,  "/" },
    { token.T_ASSIGN, "=" },
    { token.T_OPEN_PAREN,  "(" },
    { token.T_CLOSE_PAREN, ")" },
    { token.T_IDENTIFIER,  "variable" },
    { token.T_INTEGER,     "10" },
    { token.T_BANG,        "!" },
  }

  scanner := New(`+ - * / = ( ) variable 10 !`)

  for _, assertion := range assertions {
    subject := scanner.NextToken()

    if subject.Type != assertion.ExpectedType {
      t.Fatalf("Incorrect token type. expected=%q, got=%q",
          assertion.ExpectedType, subject.Type)
    }

    if subject.Literal != assertion.ExpectedLiteral {
      t.Fatalf("Incorrect literal type. expected=%q, got=%q",
          assertion.ExpectedLiteral, subject.Literal)
    }
  }
}