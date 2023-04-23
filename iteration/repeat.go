package iteration

import "strings"

func Repeat(character string, times int) string {
	var repeated = strings.Repeat(character, times)

	return repeated
}

func Replace(text, characterMatch, characterReplace string) string {
	var replaced = strings.ReplaceAll(text, characterMatch, characterReplace)

	return replaced
}

func Contains(text string, substring string) bool {
	return strings.Contains(text, substring)
}
