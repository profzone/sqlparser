package sql

func Parse(lexer *Lexer) {
	yyErrorVerbose = true
	yyParse(lexer)
}
