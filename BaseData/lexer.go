package basedata

var SeparatorLogic = []string{
	" ",
	"	",
}

func Lexer(input string) [][]string {
	var value [][]string

	currentWord := ""
	for _, char := range input {
		isSeparator := false
		for _, sep := range SeparatorLogic {
			if string(char) == sep {
				isSeparator = true
				break
			}
		}

		if isSeparator {
			if currentWord != "" {
				value = append(value, []string{currentWord})
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}

	if currentWord != "" {
		value = append(value, []string{currentWord})
	}

	return value
}

func Lexer2(input string) []string {
	var value []string

	currentWord := ""
	for _, char := range input {
		isSeparator := false
		for _, sep := range SeparatorLogic {
			if string(char) == sep {
				isSeparator = true
				break
			}
		}

		if isSeparator {
			if currentWord != "" {
				value = append(value, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}

	if currentWord != "" {
		value = append(value, currentWord)
	}

	return value
}
