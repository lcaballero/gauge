package lexer

type LexOptions struct {
	Interpolated   bool
	StartingLine   int
	StartingColumn int
	Plugins        interface{} // TODO
}

func NewLexOptions() *LexOptions {
	return &LexOptions{
		Interpolated:   false,
		StartingColumn: 1,
		StartingLine:   1,
		Plugins:        nil,
	}
}
