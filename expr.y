%{

package parser

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
)

%}

%union {
	value *Value
	num float64
}

%type <value> expr list nil atom 

%token CAR CDR CONS NIL

%token <num> NUM

%%

expr:
	list
|	atom

list:
	'(' CAR list ')'
	{
		$$ = $3.Left
	}
|	'(' CDR list ')'
	{
		$$ = $3.Right
	}
|	'(' CONS expr expr ')'
	{
		$$ = &Value{Left: $3, Right: $4}
	}
|	nil

nil:
	NIL
	{
		$$ = &Value{IsNil: true}
	}
|	'(' ')'
	{
		$$ = &Value{IsNil: true}
	}

atom: 
	NUM
	{
		$$ = &Value{Num: $1}
	}

%%
