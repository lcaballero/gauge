package lexer

import (
	"testing"

	"bytes"

	. "github.com/smartystreets/goconvey/convey"
)


func TestLexer(t *testing.T) {

	jadeBytes := []byte(`
doctype html
html
   head
      title
body
   div.className(atts='here')
`)

	Convey("Creates an initial token for 'doc-start'", t, func() {
		lex, _ := NewLexer(jadeBytes, "filename.jade", NewLexOptions())
		lex.DocStart()
		So(lex.tokens, ShouldNotBeNil)
		So(len(lex.tokens), ShouldEqual, 1)

		tk := lex.tokens[0]
		So(tk.Type, ShouldEqual, DocStart)
		So(string(tk.Val), ShouldEqual, "filename.jade")
	})

	Convey("Given 0 mark Lexer should find the doctype", t, func() {
		lex, _ := NewLexer(jadeBytes, "filename.txt", NewLexOptions())
		lex.DocStart()
		lex.DocType()
		So(lex.tokens, ShouldNotBeNil)
		So(len(lex.tokens), ShouldEqual, 2)

		So(lex.tokens[0].Type, ShouldEqual, DocStart)
		So(lex.tokens[1].Type, ShouldEqual, DocType)

		So(string(lex.tokens[0].Val), ShouldEqual, "filename.txt")
		So(string(lex.tokens[1].Val), ShouldEqual, "html")
	})

	Convey("New token should hold given type and value along with current line and column number", t, func() {
		lex, _ := NewLexer(jadeBytes, "", NewLexOptions())
		lex.IncrementLine(10)
		lex.IncrementColumn(5)
		tk := lex.Tok(Unknown, []byte("test-value"))
		So(tk.Type, ShouldEqual, Unknown)
		So(tk.Col, ShouldEqual, 6)
		So(tk.Line, ShouldEqual, 11)
		So(bytes.Equal(tk.Val, []byte("test-value")), ShouldBeTrue)
	})

	Convey("IncrementLine increases lexer line by incrment.", t, func() {
		lex, _ := NewLexer(jadeBytes, "", NewLexOptions())
		lex.IncrementColumn(10)
		So(lex.colNum, ShouldEqual, 11)
		lex.IncrementLine(1)
		So(lex.colNum, ShouldEqual, 1)
		So(lex.lineNum, ShouldEqual, 2)
	})

	Convey("IncrementLine increases lexer line by incrment.", t, func() {
		lex, _ := NewLexer(jadeBytes, "", NewLexOptions())
		lex.IncrementColumn(2)
		So(lex.colNum, ShouldEqual, 3)
	})

	Convey("New lexer cannot be created with nil LexerOptions", t, func() {
		lex, err := NewLexer(jadeBytes, "", NewLexOptions())
		So(lex, ShouldNotBeNil)
		So(err, ShouldBeNil)
	})

	Convey("New lexer cannot be created with nil LexerOptions", t, func() {
		lex, err := NewLexer(jadeBytes, "", nil)
		So(lex, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("New lexer cannot be created with empty byte slice", t, func() {
		lex, err := NewLexer([]byte{}, "", nil)
		So(lex, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})

	Convey("New lexer cannot be created with nil byte slice", t, func() {
		lex, err := NewLexer(nil, "", nil)
		So(lex, ShouldBeNil)
		So(err, ShouldNotBeNil)
	})
}
