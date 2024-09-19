package morse

import "strings"

func TextToMorse(text string) (status string, res string) {
	for _, v := range strings.ToUpper(text) {
		res += " " + Translate(v)
	}
	if len(res) < 1 {
		return "fail", res
	}
	return "success", res
}

func Translate(char rune) string {
	return Codes[string(char)]
}
