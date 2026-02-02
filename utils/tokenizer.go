package utils

func analyze(text string) []string {
	tokens := Tokenize(text)
	tokens = lowerCaseFilter(tokens)
	tokens = stopWordFilter(tokens)
	tokens = stemmerFilter(tokens)
	return tokens
}
