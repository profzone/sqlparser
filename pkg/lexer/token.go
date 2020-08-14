package lexer

func RegisterTokenFunc(tokenStart string, fn TokenScanFunc) {
	for _, c := range tokenStart {
		if Tokens.children[c] == nil {
			Tokens.children[c] = &trieNode{}
		}
		Tokens.children[c].fn = fn
	}
}

func RegisterToken(ch byte, tok int) {
	if Tokens.children[ch] == nil {
		Tokens.children[ch] = &trieNode{}
	}
	Tokens.children[ch].token = tok
}

func RegisterTokenStr(str string, tok int) {
	var node = &Tokens
	for _, c := range str {
		if node.children[c] == nil {
			node.children[c] = &trieNode{}
		}
		node = node.children[c]
	}
	node.token = tok
}
