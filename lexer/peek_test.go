package lexer_test

import (
	"testing"

	"github.com/alecthomas/assert/v2"

	"github.com/alecthomas/participle/v2/lexer"
)

type staticLexer struct {
	tokens []lexer.Token
}

func (s *staticLexer) Next() (lexer.Token, error) {
	if len(s.tokens) == 0 {
		return lexer.EOFToken(lexer.Position{}), nil
	}
	t := s.tokens[0]
	s.tokens = s.tokens[1:]
	return t, nil
}

func TestUpgrade(t *testing.T) {
	t0 := lexer.Token{Type: 1, Value: "moo"}
	ts := lexer.Token{Type: 3, Value: " "}
	t1 := lexer.Token{Type: 2, Value: "blah"}
	tokens := []lexer.Token{t0, ts, t1}
	l, err := lexer.Upgrade(&staticLexer{tokens: tokens}, 3)
	assert.NoError(t, err)
	assert.Equal(t, t0, mustPeek(t, l, 0))
	assert.Equal(t, t0, mustPeek(t, l, 0))
	assert.Equal(t, t1, mustPeek(t, l, 1))
	assert.Equal(t, t1, mustPeek(t, l, 1))
	assert.True(t, mustPeek(t, l, 2).EOF())
	assert.True(t, mustPeek(t, l, 3).EOF())
	assert.Equal(t, tokens, l.Range(0, 3))
}

func mustPeek(t *testing.T, lexer *lexer.PeekingLexer, n int) lexer.Token {
	t.Helper()
	token, err := lexer.Peek(n)
	assert.NoError(t, err)
	return token
}

func mustNext(t *testing.T, lexer lexer.Lexer) lexer.Token {
	t.Helper()
	token, err := lexer.Next()
	assert.NoError(t, err)
	return token
}
