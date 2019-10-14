package parser

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLexerLex(t *testing.T) {
	tests := []struct {
		expr []byte
		outs []int
		symTypes []*exprSymType
	} {
		{
			expr: []byte("123.456"),
			outs: []int{NUM},
			symTypes: []*exprSymType{
				{
					num: 123.456,
				},
			},
		},
		{
			expr: []byte("(cons 1 2)"),
			outs: []int{'(', CONS, NUM, NUM, ')'},
			symTypes: []*exprSymType{
				{},
				{},
				{
					num: 1,
				},
				{
					num: 2,
				},
				{},
			},
		},
		{
			expr: []byte("(cdr (CONS (car (cons () NIL)) 1e5))"),
			outs: []int{'(', CDR, '(', CONS, '(', CAR, '(', CONS, '(', ')', NIL, ')', ')', NUM, ')', ')'},
			symTypes: []*exprSymType{
				{},
				{}, // CDR
				{},
				{}, // CONS
				{},
				{}, // CAR
				{},
				{}, // CONS
				{},
				{},
				{}, // NIL
				{},
				{},
				{
					num: 1e5,
				},
				{},
				{},
			},
		},
	}
	for _, test := range tests {
		t.Run(fmt.Sprintf("%s -> %v", string(test.expr), test.outs), func(t *testing.T) {
			i := 0
			lexer := newLexer([]byte(test.expr))
			for {
				symType := &exprSymType{}
				char := lexer.Lex(symType)
				if char == eof {
					break
				}
				assert.Equal(t, test.outs[i], char)
				assert.Equal(t, test.symTypes[i], symType)
				i++
			}
			assert.Equal(t, i, len(test.outs))
		})
	}
}
