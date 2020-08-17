package common

func IsLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

func IsDigit(ch rune) bool {
	return ch >= '0' && ch <= '9'
}

func IsIdentChar(ch rune) bool {
	return IsLetter(ch) || IsDigit(ch) || ch == '_' || ch == '$' || IsIdentExtend(ch)
}

func IsIdentExtend(ch rune) bool {
	return ch >= 0x80 && ch <= '\uffff'
}

func IsUserVarChar(ch rune) bool {
	return IsLetter(ch) || IsDigit(ch) || ch == '_' || ch == '$' || ch == '.' || IsIdentExtend(ch)
}

func IsBool(b string) (bool, bool) {
	switch b {
	case "true", "TRUE", "True":
		return true, true
	case "false", "FALSE", "False":
		return false, true
	}
	return false, false
}
