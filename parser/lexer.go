package parser

import (
	"bytes"
	"errors"
	"log"
	"strconv"
	"unicode"
	"unicode/utf8"
)

type lexer struct {
	line []byte
	peek *rune
	err error
}

func newLexer(line []byte) *lexer {
	return &lexer{
		line: line,
	}
}

const eof = 0

func (x *lexer) Lex(yylval *exprSymType) int {
	for {
		c := x.next()
		if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z'  {
			return x.token(c, yylval)
		}
		switch c {
		case eof:
			return eof
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			return x.num(c, yylval)
		case '(', ')':
			return int(c)

		case ' ', '\t', '\n', '\r':
		default:
			log.Printf("unrecognized character %q", c)
		}
	}
}

func (x *lexer) next() rune {
	if x.peek != nil {
		r := *x.peek
		x.peek = nil
		return r
	}
	if len(x.line) == 0 {
		return eof
	}
	c, size := utf8.DecodeRune(x.line)
	x.line = x.line[size:]
	if c == utf8.RuneError && size == 1 {
		log.Print("invalid utf8")
		return x.next()
	}
	return c
}

func (x *lexer) num(c rune, yylval *exprSymType) int {
	add := func(b *bytes.Buffer, c rune) {
		if _, err := b.WriteRune(c); err != nil {
			log.Fatalf("WriteRune: %s", err)
		}
	}
	var b bytes.Buffer
	add(&b, c)
L:
	for {
		c = x.next()
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', 'e', 'E':
			add(&b, c)
		default:
			x.peek = &c
			break L
		}
	}
	num, err := strconv.ParseFloat(b.String(), 10)
	if err != nil {
		log.Printf("bad number %q", b.String())
		return eof
	}
	yylval.num = num
	return NUM
}

var tokenMap = map[string]int{
	"car": CAR,
	"cdr": CDR,
	"cons": CONS,
	"nil": NIL,
}

func (x *lexer) token(c rune, yylval *exprSymType) int {
	add := func(b *bytes.Buffer, c rune) {
		if _, err := b.WriteRune(c); err != nil {
			log.Fatalf("WriteRune: %s", err)
		}
	}
	var b bytes.Buffer
	add(&b, unicode.ToLower(c))
	for {
		c = x.next()
		if 'a' <= c && c <= 'z' || 'A' <= c && c <= 'Z' {
			add(&b, unicode.ToLower(c))
		} else {
			x.peek = &c
			break
		}
	}
	token, ok := tokenMap[b.String()]
	if !ok {
		log.Printf("unknown token %q", b.String())
		return eof
	}
	return token
}


func (x *lexer) Error(s string) {
	x.err = errors.New(s)
}
