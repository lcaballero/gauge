package lexer

import (
	"errors"
	"regexp"
	"fmt"
)

// Lexer tokenizes Jade text into Tokens instances.
type Lexer struct {
	input                []byte
	originalInput        []byte
	filename             string
	interpolated         bool
	lineNum              int
	colNum               int
	plugins              []interface{} // TODO: plugins
	indentStack          []int
	indentRe             string
	interpolationAllowed bool
	tokens               []*Token
	ended                bool
	opts                 *LexOptions

	mark int // current read position
}

// NewLexer creates a new lexer over the given input for the filename
// provided and with the LexOptions.
func NewLexer(input []byte, filename string, opts *LexOptions) (*Lexer, error) {
	if opts == nil {
		return nil, errors.New("LexOptions cannot be nil.")
	}
	if input == nil || len(input) <= 0 {
		return nil, errors.New("input byte slice cannot be nil or empty.")
	}
	lex := &Lexer{
		originalInput:        input,
		input:                input,
		filename:             filename,
		ended:                false,
		interpolationAllowed: false,
		interpolated:         false,
		indentStack:          []int{0},
		colNum:               opts.StartingColumn,
		lineNum:              opts.StartingLine,
		opts:                 opts,
		mark:                 0,
		tokens:               make([]*Token, 0),
	}
	return lex, nil
}

// Tok creates a new Token with the given type and string value.
func (x *Lexer) Tok(ttype TokenType, val []byte) *Token {
	return &Token{
		Type: ttype,
		Line: x.lineNum,
		Col:  x.colNum,
		Val:  val,
	}
}

// Emmit creates a new Token and saves that Token internally.
func (x *Lexer) Emit(ttype TokenType, val []byte) {
	x.tokens = append(x.tokens, x.Tok(ttype, val))
}

// IncrementLine increases the line number by n and resets column
// to start of line.
func (x *Lexer) IncrementLine(n int) {
	x.lineNum += n
	x.colNum = x.opts.StartingColumn
}

// IncrementColumn increases the column by n.
func (x *Lexer) IncrementColumn(n int) {
	x.colNum += n
}

// MatchBounds extracts start and end from regex Indexes.
func (x *Lexer) MatchBounds(indexes [][]int) (int, int) {
	return indexes[0][0], indexes[0][1]
}

// DocStart consumes any empty strings prior to start of parsing document.
func (x *Lexer) DocStart() {
	re := regexp.MustCompile("^[ \n]*")
	indexes := re.FindAllIndex(x.input, 1)

	if len(indexes) <= 0 {
		return
	}

	fmt.Println("DocStart", indexes)

	_, end := x.MatchBounds(indexes)
	x.Emit(DocStart, []byte(x.filename))
	x.input = x.input[end:]
}

// DocType consumes doctype if detected and emits a DocType token.
// /^doctype *([^\n]*)/
func (x *Lexer) DocType() {
	re := regexp.MustCompile("^doctype *([^\n]*)")
	indexes := re.FindAllIndex(x.input, 1)

	if len(indexes) <= 0 {
		return
	}

	start, end := x.MatchBounds(indexes)
	matchedBytes := x.input[start:end]
	subMatched := re.FindSubmatch(matchedBytes)

	if len(subMatched) > 1 {
		x.Emit(DocType, subMatched[1])
		x.input = x.input[end:]
	}
}
