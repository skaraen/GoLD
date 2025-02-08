package token

type Token struct {
	Line 	int
	Column 	int
}

func NewToken(line, column int) *Token {
	return &Token{Line: line, Column: column}
}