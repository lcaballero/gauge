package lexer

type TokenType int

const (
	Unknown  TokenType = 0
	DocStart TokenType = 1
	DocType  TokenType = 2
)

// {type: type, line: this.lineno, col: this.colno};
type Token struct {
	Type TokenType
	Line int
	Col  int
	Val  []byte
}
