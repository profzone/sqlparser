package lexer

import (
	"unicode"
	"unicode/utf8"
)

type Pos struct {
	Line   int
	Col    int
	Offset int
}

type lexerReader struct {
	code        string
	pos         Pos
	tokenLength int
}

func (r *lexerReader) reset(code string) {
	r.code = code
	r.pos = Pos{}
	r.tokenLength = 0
}

func (r *lexerReader) eof() bool {
	return r.pos.Offset >= len(r.code)
}

func (r *lexerReader) peek() rune {
	if r.eof() {
		return unicode.ReplacementChar
	}
	val, length := rune(r.code[r.pos.Offset]), 1
	switch {
	case val == 0:
		r.tokenLength = length
		return val // illegal UTF-8 encoding
	case val >= 0x80:
		val, length = utf8.DecodeRuneInString(r.code[r.pos.Offset:])
		if val == utf8.RuneError && length == 1 {
			val = rune(r.code[r.pos.Offset]) // illegal UTF-8 encoding
		}
	}
	r.tokenLength = length
	return val
}

func (r *lexerReader) inc() {
	if r.code[r.pos.Offset] == '\n' {
		r.pos.Line++
		r.pos.Col = 0
	}
	r.pos.Offset += r.tokenLength
	r.pos.Col++
}

func (r *lexerReader) incN(n int) {
	for i := 0; i < n; i++ {
		r.inc()
	}
}

func (r *lexerReader) data(from *Pos) string {
	return r.code[from.Offset:r.pos.Offset]
}

func (r *lexerReader) incAsLongAs(fn func(rune) bool) rune {
	for {
		ch := r.peek()
		if !fn(ch) {
			return ch
		}
		if ch == unicode.ReplacementChar && r.eof() {
			return 0
		}
		r.inc()
	}
}

type TokenScanFunc func(scanner *Scanner) (tok int, pos Pos, lit string)

var Tokens = trieNode{}

type trieNode struct {
	children [256]*trieNode
	token    int
	fn       TokenScanFunc
}

type Scanner struct {
	reader lexerReader
}

func NewScanner(code string) *Scanner {
	return &Scanner{
		reader: lexerReader{
			code: code,
		},
	}
}

func (s *Scanner) skipWhitespace() rune {
	return s.reader.incAsLongAs(unicode.IsSpace)
}

func (s *Scanner) Reset(code string) {
	s.reader.reset(code)
}

func (s *Scanner) Scan() (tok int, pos Pos, lit string) {
	ch0 := s.reader.peek()
	if unicode.IsSpace(ch0) {
		ch0 = s.skipWhitespace()
	}

	pos = s.reader.pos
	if s.reader.eof() {
		// when scanner meets EOF, the returned token should be 0,
		// because 0 is a special token id to remind the parser that stream is end.
		return 0, pos, ""
	}

	var node = &Tokens
	for ch0 >= 0 && ch0 <= 255 {
		if node.children[ch0] == nil || s.reader.eof() {
			break
		}
		node = node.children[ch0]
		if node.fn != nil {
			return node.fn(s)
		}

		s.reader.inc()
		ch0 = s.reader.peek()
	}

	tok, lit = node.token, s.reader.data(&pos)
	return
}

func (s *Scanner) Inc() {
	s.reader.inc()
}

func (s *Scanner) IncAsLongAs(fn func(rune) bool) rune {
	return s.reader.incAsLongAs(fn)
}

func (s *Scanner) Data(from *Pos) string {
	return s.reader.data(from)
}

func (s *Scanner) Pos() Pos {
	return s.reader.pos
}

func (s *Scanner) Eof() bool {
	return s.reader.eof()
}

func (s *Scanner) Peek() rune {
	return s.reader.peek()
}
