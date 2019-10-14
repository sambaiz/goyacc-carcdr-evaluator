package parser

import (
	"fmt"
	"strconv"
)

type Value struct {
	IsNil bool
	Num   float64
	Left  *Value
	Right *Value
}

func (v *Value) String() string {
	if v.IsNil {
		return "NIL"
	}
	if v.Left != nil && v.Right != nil {
		return fmt.Sprintf("(%s . %s)", v.Left.String(), v.Right.String())
	}
	return fmt.Sprintf(strconv.FormatFloat(v.Num, 'g', 12, 64))
}
