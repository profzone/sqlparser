package calc

func Parse(lexer *Lexer) {
	yyErrorVerbose = true
	yyParse(lexer)
}
