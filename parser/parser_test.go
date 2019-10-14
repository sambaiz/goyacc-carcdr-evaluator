package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		in string
		expected string
	} {
		{
			in: "123.456",
			expected: "123.456",
		},
		{
			in: "(cdr (CONS (car (cons () NIL)) 1e5))",
			expected: "100000",
		},
		{
			in: "(cons (cons 1 2) 3)",
			expected: "((1 . 2) . 3)",
		},
	}
	for _, test := range tests {
		t.Run(test.in, func(t *testing.T) {
			out, err := Parse(test.in)
			assert.Equal(t, test.expected, out.String())
			assert.Nil(t, err)
		})
	}
}