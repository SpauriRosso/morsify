package utils

import (
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
)

func VerifyPrompt(input string) (status string, result string) {
	result = NormalizePrompt(input)

	for _, v := range result {
		if !unicode.In(v, unicode.Latin) {
			return "fail : non-latin character found", ""
		}
	}
	return status, result
}

func NormalizePrompt(input string) string {
	// Normalize input by décomposing it into its canonical form (unicode/norm) & separate characters into their base charater and diatrical mark
	// example : è would become e` as the mark itself is considered a character
	text := norm.NFD.String(input)

	prompt := strings.Builder{}
	for _, v := range text {
		if unicode.IsLetter(v) && !unicode.Is(unicode.Mn, v) { // Check if v is diacritical mark
			prompt.WriteRune(v)
		}
	}
	return prompt.String()
}
