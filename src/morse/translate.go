package morse

import "strings"

func TextToMorse(text string) string {
	var res string
	for _, v := range strings.ToUpper(text) {
		res += " " + Translate(v)
	}
	return res
}

func Translate(char rune) string {
	return MorseCodes[string(char)]
}
