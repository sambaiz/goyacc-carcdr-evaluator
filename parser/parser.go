package parser

func Parse(line string) (*Value, error) {
	exprErrorVerbose = true
	parser := &exprParserImpl{}
	lexer := newLexer([]byte(line))
	parser.Parse(lexer)
	if lexer.err != nil {
		return nil, lexer.err
	}
	if len(parser.stack) < 2 {
		return nil, nil
	}
	return parser.stack[1].value, nil
}